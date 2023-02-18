// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fileconsumer // import "github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/fileconsumer"

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/pkg/stanza/operator"
)

type EmitFunc func(ctx context.Context, attrs *FileAttributes, token []byte)

type Manager struct {
	*zap.SugaredLogger
	wg     sync.WaitGroup
	cancel context.CancelFunc

	readerFactory readerFactory
	finder        Finder
	roller        roller
	persister     operator.Persister

	pollInterval    time.Duration
	maxBatchFiles   int
	deleteAfterRead bool

	knownFiles     []*Reader
	seenPaths      map[string]struct{}
	currentFiles   []*os.File
	currentFps     []*Fingerprint
	currentReaders []*Reader
}

func (m *Manager) Start(persister operator.Persister) error {
	ctx, cancel := context.WithCancel(context.Background())
	m.cancel = cancel
	m.persister = persister

	// Load offsets from disk
	if err := m.loadLastPollFiles(ctx); err != nil {
		return fmt.Errorf("read known files from database: %w", err)
	}

	if len(m.finder.FindFiles()) == 0 {
		m.Warnw("no files match the configured include patterns",
			"include", m.finder.Include,
			"exclude", m.finder.Exclude)
	}

	// Start polling goroutine
	m.startPoller(ctx)

	return nil
}

// Stop will stop the file monitoring process
func (m *Manager) Stop() error {
	m.cancel()
	m.wg.Wait()
	m.roller.cleanup()
	for _, reader := range m.knownFiles {
		reader.Close()
	}
	m.knownFiles = nil
	m.cancel = nil
	return nil
}

// startPoller kicks off a goroutine that will poll the filesystem periodically,
// checking if there are new files or new logs in the watched files
func (m *Manager) startPoller(ctx context.Context) {
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		globTicker := time.NewTicker(m.pollInterval)
		defer globTicker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-globTicker.C:
			}

			m.poll(ctx)
			m.clearCurrentFiles()
		}
	}()
}

// poll checks all the watched paths for new entries
func (m *Manager) poll(ctx context.Context) {
	// Increment the generation on all known readers
	// This is done here because the next generation is about to start
	for i := 0; i < len(m.knownFiles); i++ {
		m.knownFiles[i].generation++
	}

	// Get the list of paths on disk
	matches := m.finder.FindFiles()
	m.consume(ctx, matches)
}

func (m *Manager) consume(ctx context.Context, paths []string) {
	pathChan := make(chan string, m.maxBatchFiles)
	closeChan := make(chan int)
	pathsConsumedChan := make(chan int)
	readerChan := make(chan *Reader)
	readers := make([]*Reader, 0, len(paths))
	var wg sync.WaitGroup
	for i := 0; i < m.maxBatchFiles; i++ {
		wg.Add(1)
		go m.consumeAsync(ctx, pathsConsumedChan, pathChan, closeChan, wg)
	}
	for i := 0; i < len(paths); i++ {
		pathChan <- paths[i]
	}

	totalPathsConsumed := 0
	for i := range pathsConsumedChan {
		if totalPathsConsumed == len(paths) {
			closeChan <- 1
			break;
		}
		totalPathsConsumed += i
	}

	for r := range readerChan {
		readers = append(readers, r)
		if !m.deleteAfterRead && (len(readers) % m.maxBatchFiles == 0) {
			m.roller.readLostFiles(ctx, readers)
			m.roller.roll(ctx, readers)
			m.saveCurrent(readers)
			m.syncLastPollFiles(ctx)
			readers = make([]*Readers, 0)
		}
	}
	if  !m.deleteAfterRead && (len(readers) > 0) {
		m.roller.readLostFiles(ctx, readers)
		m.roller.roll(ctx, readers)
		m.saveCurrent(readers)
		m.syncLastPollFiles(ctx)
	}
	wg.Wait()
	close(readerChan)
	close(pathsConsumedChan)
	close(closeChan)
	close(pathChan)
}

func (m *Manager) consumeAsync(ctx context.Context, pathsConsumedChan chan int, pathChan chan string, closeChan chan int, readerChan chan *Reader, wg sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case path := <-pathChan:
			pathsConsumedChan <- 1
			r := m.makeReader(path)
			r.ReadToEnd(ctx)
			if m.deleteAfterRead {
				r.Close()
				if err := os.Remove(r.file.Name()); err != nil {
					m.Errorf("could not delete %s", r.file.Name())
				}
			} else {
				readerChan <- r
			}

		case <-closeChan:
			break
		}
	}
	return
}

// func (m *Manager) consume(ctx context.Context, paths []string) {
// 	m.Debug("Consuming files")
// 	readers := m.makeReaders(paths)

// 	// take care of files which disappeared from the pattern since the last poll cycle
// 	// this can mean either files which were removed, or rotated into a name not matching the pattern
// 	// we do this before reading existing files to ensure we emit older log lines before newer ones
// 	m.roller.readLostFiles(ctx, readers)

// 	var wg sync.WaitGroup
// 	for _, reader := range readers {
// 		wg.Add(1)
// 		go func(r *Reader) {
// 			defer wg.Done()
// 			r.ReadToEnd(ctx)
// 			if m.deleteAfterRead {
// 				r.Close()
// 				if err := os.Remove(r.file.Name()); err != nil {
// 					m.Errorf("could not delete %s", r.file.Name())
// 				}
// 			}
// 		}(reader)
// 	}
// 	wg.Wait()

// 	if m.deleteAfterRead {
// 		// no need to track files since they were deleted
// 		return
// 	}

// 	// Any new files that appear should be consumed entirely
// 	m.readerFactory.fromBeginning = true

// 	m.roller.roll(ctx, readers)
// 	m.saveCurrent(readers)
// 	m.syncLastPollFiles(ctx)
// }

// makeReaders takes a list of paths, then creates readers from each of those paths,
// discarding any that have a duplicate fingerprint to other files that have already
// been read this polling interval
// func (m *Manager) makeReaders(filesPaths []string) []*Reader {
// 	// Open the files first to minimize the time between listing and opening
// 	files := make([]*os.File, 0, len(filesPaths))
// 	for _, path := range filesPaths {
// 		if _, ok := m.seenPaths[path]; !ok {
// 			if m.readerFactory.fromBeginning {
// 				m.Infow("Started watching file", "path", path)
// 			} else {
// 				m.Infow("Started watching file from end. To read preexisting logs, configure the argument 'start_at' to 'beginning'", "path", path)
// 			}
// 			m.seenPaths[path] = struct{}{}
// 		}
// 		file, err := os.Open(path) // #nosec - operator must read in files defined by user
// 		if err != nil {
// 			m.Debugf("Failed to open file", zap.Error(err))
// 			continue
// 		}
// 		files = append(files, file)
// 	}

// 	// Get fingerprints for each file
// 	fps := make([]*Fingerprint, 0, len(files))
// 	for _, file := range files {
// 		fp, err := m.readerFactory.newFingerprint(file)
// 		if err != nil {
// 			m.Errorw("Failed creating fingerprint", zap.Error(err))
// 			continue
// 		}
// 		fps = append(fps, fp)
// 	}

// 	// Exclude any empty fingerprints or duplicate fingerprints to avoid doubling up on copy-truncate files
// OUTER:
// 	for i := 0; i < len(fps); i++ {
// 		fp := fps[i]
// 		if len(fp.FirstBytes) == 0 {
// 			if err := files[i].Close(); err != nil {
// 				m.Errorf("problem closing file", "file", files[i].Name())
// 			}
// 			// Empty file, don't read it until we can compare its fingerprint
// 			fps = append(fps[:i], fps[i+1:]...)
// 			files = append(files[:i], files[i+1:]...)
// 			i--
// 			continue
// 		}
// 		for j := i + 1; j < len(fps); j++ {
// 			fp2 := fps[j]
// 			if fp.StartsWith(fp2) || fp2.StartsWith(fp) {
// 				// Exclude
// 				if err := files[i].Close(); err != nil {
// 					m.Errorf("problem closing file", "file", files[i].Name())
// 				}
// 				fps = append(fps[:i], fps[i+1:]...)
// 				files = append(files[:i], files[i+1:]...)
// 				i--
// 				continue OUTER
// 			}
// 		}
// 	}

// 	readers := make([]*Reader, 0, len(fps))
// 	for i := 0; i < len(fps); i++ {
// 		reader, err := m.newReader(files[i], fps[i])
// 		if err != nil {
// 			m.Errorw("Failed to create reader", zap.Error(err))
// 			continue
// 		}
// 		readers = append(readers, reader)
// 	}

// 	return readers
// }

func (m *Manager) makeReaders(filePaths []string) []*Reader {
	readers := make([]*Readers, 0)
	for i := 0; i < len(filePaths); i++ {
		reader := m.makeReader(path)
		if reader == nil {
			m.Errorw("Failed to create reader for path ", path)
			continue
		}
		readers = append(readers, reader)
	}

	return readers
}

// makeReader takes a path and it return reader for that path
// discarding any that have a duplicate fingerprint to other files that have already
// been read this polling interval
func (m *Manager) makeReader(filePath string) *Reader {
	if _, ok := m.seenPaths[filePath]; !ok {
		if m.readerFactory.fromBeginning {
			m.Infow("Started watching file", "path", filePath)
		} else {
			m.Infow("Startedd watching file from end. To read preexisting logs, configure the argument 'start_at' to 'beginning'", "path", path)
		}
		m.seenPaths[filePath] = struct{}{}
	}
	file, err := os.Open(filePath) // #nosec - operator must read in files defined by user
	if err != nil {
		m.Debugf("Failed to open file", zap.Error(err))
		return nil
	}
	m.currentFiles = append(m.currentFiles, file)
	fp, err := m.readerFactory.newFingerprint(file)
	if err != nil {
		m.Errorw("Failed creating fingerprint", zap.Error(err))
		return nil
	}
	if len(fp.FirstBytes) == 0 {
		if err := file.Close(); err != nil {
			m.Errorf("problem closing file", "file", file.Name())
		}
		// Empty file, don't read it until we can compare its fingerprint
		m.currentFiles = m.currentFiles[:len(m.currentFiles)-1]
		return nil
	}
	m.currentFps = append(m.currentFps, fp)

	for i := 0; i < len(m.currentFps)-1; i++ {
		fp2 := m.currentFps[i]
		if fp.StartsWith(fp2) || fp2.StartsWith(fp) {
			// Exclude
			if err := file.Close(); err != nil {
				m.Errorf("problem closing file", "file", file.Name())
			}
			m.currentFiles = m.currentFiles[:len(m.currentFiles)-1]
			m.currentFps = m.currentFps[:len(m.currentFps)-1]
			i--
			return nil
		}
	}

	reader, err := m.newReader(file, fp)
	if err != nil {
		m.Errorw("Failed to create reader", zap.Error(err))
		return nil
	}
	return reader
}

func (m *Manager) clearCurrentFiles() {
	m.currentFiles = make([]*os.File, 0)
	m.currentFps = make([]*Fingerprint, 0)
}

// saveCurrent adds the readers from this polling interval to this list of
// known files, then increments the generation of all tracked old readers
// before clearing out readers that have existed for 3 generations.
func (m *Manager) saveCurrent(readers []*Reader) {
	// Add readers from the current, completed poll interval to the list of known files
	m.knownFiles = append(m.knownFiles, readers...)

	// Clear out old readers. They are sorted such that they are oldest first,
	// so we can just find the first reader whose generation is less than our
	// max, and keep every reader after that
	for i := 0; i < len(m.knownFiles); i++ {
		reader := m.knownFiles[i]
		if reader.generation <= 3 {
			m.knownFiles = m.knownFiles[i:]
			break
		}
	}
}

func (m *Manager) newReader(file *os.File, fp *Fingerprint) (*Reader, error) {
	// Check if the new path has the same fingerprint as an old path
	if oldReader, ok := m.findFingerprintMatch(fp); ok {
		return m.readerFactory.copy(oldReader, file)
	}

	// If we don't match any previously known files, create a new reader from scratch
	return m.readerFactory.newReader(file, fp)
}

func (m *Manager) findFingerprintMatch(fp *Fingerprint) (*Reader, bool) {
	// Iterate backwards to match newest first
	for i := len(m.knownFiles) - 1; i >= 0; i-- {
		oldReader := m.knownFiles[i]
		if fp.StartsWith(oldReader.Fingerprint) {
			return oldReader, true
		}
	}
	return nil, false
}

const knownFilesKey = "knownFiles"

// syncLastPollFiles syncs the most recent set of files to the database
func (m *Manager) syncLastPollFiles(ctx context.Context) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)

	// Encode the number of known files
	if err := enc.Encode(len(m.knownFiles)); err != nil {
		m.Errorw("Failed to encode known files", zap.Error(err))
		return
	}

	// Encode each known file
	for _, fileReader := range m.knownFiles {
		if err := enc.Encode(fileReader); err != nil {
			m.Errorw("Failed to encode known files", zap.Error(err))
		}
	}

	if err := m.persister.Set(ctx, knownFilesKey, buf.Bytes()); err != nil {
		m.Errorw("Failed to sync to database", zap.Error(err))
	}
}

// syncLastPollFiles loads the most recent set of files to the database
func (m *Manager) loadLastPollFiles(ctx context.Context) error {
	encoded, err := m.persister.Get(ctx, knownFilesKey)
	if err != nil {
		return err
	}

	if encoded == nil {
		m.knownFiles = make([]*Reader, 0, 10)
		return nil
	}

	dec := json.NewDecoder(bytes.NewReader(encoded))

	// Decode the number of entries
	var knownFileCount int
	if err := dec.Decode(&knownFileCount); err != nil {
		return fmt.Errorf("decoding file count: %w", err)
	}

	if knownFileCount > 0 {
		m.Infow("Resuming from previously known offset(s). 'start_at' setting is not applicable.")
		m.readerFactory.fromBeginning = true
	}

	// Decode each of the known files
	m.knownFiles = make([]*Reader, 0, knownFileCount)
	for i := 0; i < knownFileCount; i++ {
		// Only the offset, fingerprint, and splitter
		// will be used before this reader is discarded
		unsafeReader, err := m.readerFactory.unsafeReader()
		if err != nil {
			return err
		}
		if err = dec.Decode(unsafeReader); err != nil {
			return err
		}
		m.knownFiles = append(m.knownFiles, unsafeReader)
	}

	return nil
}
