// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package cache

import "github.com/gittuf/gittuf/pkg/gitinterface"

func (p *Persistent) SetLastVerifiedEntryForRef(ref string, entryNumber uint64, entryID gitinterface.Hash) {
	_ = "STUB: not implemented"
	return
}

// If set verified number is higher than entryNumber, noop

func (p *Persistent) GetLastVerifiedEntryForRef(ref string) (uint64, gitinterface.Hash) {
	_ = "STUB: not implemented"
	return 0, *new(gitinterface.Hash)
}
