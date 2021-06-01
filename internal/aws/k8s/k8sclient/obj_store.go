// Copyright  OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8sclient

import (
	"fmt"
	"sync"

	"go.uber.org/zap"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/types"
)

// ObjStore implements the cache.Store interface:
// https://github.com/kubernetes/client-go/blob/release-1.20/tools/cache/store.go#L26-L71
// It is used by cache.Reflector to keep the updated information about resources
// https://github.com/kubernetes/client-go/blob/release-1.20/tools/cache/reflector.go#L48
type ObjStore struct {
	mu sync.RWMutex

	refreshed bool
	objs      map[types.UID]interface{}

	transformFunc func(interface{}) (interface{}, error)
	logger        *zap.Logger
}

func NewObjStore(transformFunc func(interface{}) (interface{}, error), logger *zap.Logger) *ObjStore {
	return &ObjStore{
		transformFunc: transformFunc,
		objs:          map[types.UID]interface{}{},
		logger:        logger,
	}
}

// GetResetRefreshStatus tracks whether the underlying data store is refreshed or not.
// Calling this func itself will reset the state to false.
func (s *ObjStore) GetResetRefreshStatus() bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	refreshed := s.refreshed
	if refreshed {
		s.refreshed = false
	}
	return refreshed
}

// Add implements the Add method of the store interface.
// Add adds an entry to the ObjStore.
func (s *ObjStore) Add(obj interface{}) error {
	o, err := meta.Accessor(obj)
	if err != nil {
		s.logger.Warn(fmt.Sprintf("Cannot find the metadata for %v.", obj))
		return err
	}

	var toCacheObj interface{}
	if toCacheObj, err = s.transformFunc(obj); err != nil {
		s.logger.Warn(fmt.Sprintf("Failed to update obj %v in the cached store.", obj))
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.objs[o.GetUID()] = toCacheObj
	s.refreshed = true

	return nil
}

// Update implements the Update method of the store interface.
// Update updates the existing entry in the ObjStore.
func (s *ObjStore) Update(obj interface{}) error {
	return s.Add(obj)
}

// Delete implements the Delete method of the store interface.
// Delete deletes an existing entry in the ObjStore.
func (s *ObjStore) Delete(obj interface{}) error {

	o, err := meta.Accessor(obj)
	if err != nil {
		return err
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.objs, o.GetUID())

	s.refreshed = true

	return nil
}

// List implements the List method of the store interface.
// List lists all the objects in the ObjStore
func (s *ObjStore) List() []interface{} {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]interface{}, 0, len(s.objs))
	for _, v := range s.objs {
		result = append(result, v)
	}
	return result
}

// ListKeys implements the ListKeys method of the store interface.
// ListKeys lists the keys for all objects in the ObjStore
func (s *ObjStore) ListKeys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]string, 0, len(s.objs))
	for k := range s.objs {
		result = append(result, string(k))
	}
	return result
}

// Get implements the Get method of the store interface.
func (s *ObjStore) Get(obj interface{}) (item interface{}, exists bool, err error) {
	return nil, false, nil
}

// GetByKey implements the GetByKey method of the store interface.
func (s *ObjStore) GetByKey(key string) (item interface{}, exists bool, err error) {
	return nil, false, nil
}

// Replace implements the Replace method of the store interface.
// Replace will delete the contents of the store, using instead the given list.
func (s *ObjStore) Replace(list []interface{}, _ string) error {
	s.mu.Lock()
	s.objs = map[types.UID]interface{}{}
	s.mu.Unlock()

	for _, o := range list {
		err := s.Add(o)
		if err != nil {
			return err
		}
	}

	return nil
}

// Resync implements the Resync method of the store interface.
func (s *ObjStore) Resync() error {
	return nil
}
