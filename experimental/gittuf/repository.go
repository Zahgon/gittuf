// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gittuf

import (
	"errors"

	"github.com/gittuf/gittuf/internal/tuf"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

const (
	DebugModeKey = "GITTUF_DEBUG"
)

var (
	ErrUnauthorizedKey    = errors.New("unauthorized key presented when updating gittuf metadata")
	ErrCannotReinitialize = errors.New("cannot reinitialize metadata, it exists already")
)

// InDebugMode returns true if gittuf is currently in debug mode.
func InDebugMode() bool { _ = "STUB: not implemented"; return false }

type Repository struct {
	r *gitinterface.Repository
}

func (r *Repository) GetGitRepository() *gitinterface.Repository {
	_ = "STUB: not implemented"
	return nil
}

func LoadRepository(repositoryPath string) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func isKeyAuthorized(authorizedKeyIDs []tuf.Principal, keyID string) bool {
	_ = "STUB: not implemented"
	return false
}
