// Copyright  The OpenTelemetry Authors
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

package cache

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCache(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create a new Cache with length 10",
			args: args{
				size: 10,
			},
			wantErr: false,
		},
		{
			name: "create a new Cache with length -1",
			args: args{
				size: -1,
			},
			wantErr: true,
		},
		{
			name: "create a new Cache with length 0",
			args: args{
				size: 0,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, err := NewCache(tt.args.size)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func TestCache_Get(t *testing.T) {
	tests := []struct {
		name         string
		lruCache     func() *Cache
		evictedItems map[interface{}]interface{}
		key          interface{}
		wantValue    interface{}
		wantOk       bool
	}{
		{
			name: "if key is not found in LRUCache, will get key from evictedItems",
			lruCache: func() *Cache {
				cache, _ := NewCache(1)
				cache.evictedItems["key"] = "val"
				return cache
			},
			key:       "key",
			wantValue: "val",
			wantOk:    true,
		},
		{
			name: "if key is found in LRUCache, return the found item",
			lruCache: func() *Cache {
				cache, _ := NewCache(1)
				cache.Add("key", "val_from_LRU")
				cache.evictedItems["key"] = "val_from_evicted_items"
				return cache
			},
			key:       "key",
			wantValue: "val_from_LRU",
			wantOk:    true,
		},
		{
			name: "if key is not found either in LRUCache or evicted items, return nothing",
			lruCache: func() *Cache {
				cache, _ := NewCache(1)
				return cache
			},
			key:       "key",
			wantValue: nil,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := tt.lruCache()
			gotValue, gotOk := c.Get(tt.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestCache_RemoveEvictedItems(t *testing.T) {
	tests := []struct {
		name     string
		lruCache func() *Cache
	}{
		{
			name: "no panic when there is no evicted item to remove",
			lruCache: func() *Cache {
				cache, _ := NewCache(1)
				return cache
			},
		},
		{
			name: "evicted items should be removed",
			lruCache: func() *Cache {
				cache, _ := NewCache(1)
				cache.evictedItems["key0"] = "val0"
				cache.evictedItems["key1"] = "val1"
				return cache
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			cache := tt.lruCache()
			cache.RemoveEvictedItems()
			assert.Empty(t, cache.evictedItems)
		})
	}
}
