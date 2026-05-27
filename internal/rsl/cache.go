// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package rsl

import (
	"sync"

	"github.com/gittuf/gittuf/pkg/gitinterface"
)

type rslCache struct {
	entryCache  map[string]Entry
	parentCache map[string]string

	entryCacheMutex  sync.RWMutex
	parentCacheMutex sync.RWMutex
}

func (r *rslCache) getEntry(id gitinterface.Hash) (Entry, bool) {
	_ = "STUB: not implemented"
	return *new(Entry), false
}

func (r *rslCache) setEntry(id gitinterface.Hash, entry Entry) { _ = "STUB: not implemented"; return }

func (r *rslCache) getParent(id gitinterface.Hash) (gitinterface.Hash, bool, error) {
	_ = "STUB: not implemented"
	return *new(gitinterface.Hash), false, nil
}

func (r *rslCache) setParent(id, parentID gitinterface.Hash) { _ = "STUB: not implemented"; return }

var cache *rslCache

func newRSLCache() { _ = "STUB: not implemented"; return }

func init() {
	newRSLCache()
}
