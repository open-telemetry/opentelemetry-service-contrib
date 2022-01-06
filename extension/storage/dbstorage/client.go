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

package dbstorage // import "github.com/open-telemetry/opentelemetry-collector-contrib/extension/storage/dbstorage"

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/mattn/go-sqlite3"
	"go.opentelemetry.io/collector/extension/experimental/storage"
)

const (
	createTable     = "create table if not exists %s (key text primary key, value blob)"
	getQueryText    = "select value from %s where key=?"
	setQueryText    = "insert into %s(key, value) values(?,?) on conflict(key) do update set value=?"
	deleteQueryText = "delete from %s where key=?"
)

type dbStorageClient struct {
	db          *sql.DB
	getQuery    *sql.Stmt
	setQuery    *sql.Stmt
	deleteQuery *sql.Stmt
}

func newClient(ctx context.Context, db *sql.DB, tableName string) (*dbStorageClient, error) {
	_, err := db.ExecContext(ctx, fmt.Sprintf(createTable, tableName))
	if err != nil {
		return nil, err
	}

	selectQuery, err := db.PrepareContext(ctx, fmt.Sprintf(getQueryText, tableName))
	setQuery, err := db.PrepareContext(ctx, fmt.Sprintf(setQueryText, tableName))
	deleteQuery, err := db.PrepareContext(ctx, fmt.Sprintf(deleteQueryText, tableName))
	return &dbStorageClient{db, selectQuery, setQuery, deleteQuery}, nil
}

// Get will retrieve data from storage that corresponds to the specified key
func (c *dbStorageClient) Get(ctx context.Context, key string) ([]byte, error) {
	rows, err := c.getQuery.QueryContext(ctx, key)
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	var result []byte
	err = rows.Scan(&result)
	if err != nil {
		return result, err
	}
	err = rows.Close()
	return result, err
}

// Set will store data. The data can be retrieved using the same key
func (c *dbStorageClient) Set(ctx context.Context, key string, value []byte) error {
	_, err := c.setQuery.ExecContext(ctx, key, value, value)
	return err
}

// Delete will delete data associated with the specified key
func (c *dbStorageClient) Delete(ctx context.Context, key string) error {
	_, err := c.deleteQuery.ExecContext(ctx, key)
	return err
}

// Batch executes the specified operations in order. Get operation results are updated in place
func (c *dbStorageClient) Batch(ctx context.Context, ops ...storage.Operation) error {
	var err error
	for _, op := range ops {
		switch op.Type {
		case storage.Get:
			op.Value, err = c.Get(ctx, op.Key)
		case storage.Set:
			err = c.Set(ctx, op.Key, op.Value)
		case storage.Delete:
			err = c.Delete(ctx, op.Key)
		default:
			return errors.New("wrong operation type")
		}

		if err != nil {
			return err
		}
	}
	return err
}

// Close will close the database
func (c *dbStorageClient) Close(_ context.Context) error {
	if err := c.setQuery.Close(); err != nil {
		return err
	}
	if err := c.deleteQuery.Close(); err != nil {
		return err
	}
	if err := c.getQuery.Close(); err != nil {
		return err
	}
	return nil
}
