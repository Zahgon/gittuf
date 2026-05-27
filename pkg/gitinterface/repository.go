// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"bytes"
	"errors"
	"io"

	"github.com/go-git/go-git/v5"
	"github.com/jonboulle/clockwork"
)

const (
	binary           = "git"
	committerTimeKey = "GIT_COMMITTER_DATE"
	authorTimeKey    = "GIT_AUTHOR_DATE"
)

var ErrRepositoryPathNotSpecified = errors.New("repository path not specified")

// Repository is a lightweight wrapper around a Git repository. It stores the
// location of the repository's GIT_DIR.
type Repository struct {
	gitDirPath string
	clock      clockwork.Clock
}

// GetGoGitRepository returns the go-git representation of a repository. We use
// this in certain signing and verifying workflows.
func (r *Repository) GetGoGitRepository() (*git.Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetGitDir returns the GIT_DIR path for the repository.
func (r *Repository) GetGitDir() string { _ = "STUB: not implemented"; return "" }

// IsBare returns true if the repository is a bare repository.
func (r *Repository) IsBare() bool {
	_ = "STUB: not implemented"
	// TODO: this may not work when the repo is cloned with GIT_DIR set
	// elsewhere. We don't support this at the moment, so it's probably okay?
	return false
}

// LoadRepository returns a Repository instance using the current working
// directory. It also inspects the PATH to ensure Git is installed.
func LoadRepository(repositoryPath string) (*Repository, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

// git rev-parse --git-dir returns a local path, so filepath.Abs gives us
// the final path _including_ symlink follows.

// executor is a lightweight wrapper around exec.Cmd to run Git commands. It
// accepts the arguments to the `git` binary, but the binary itself must not be
// specified.
type executor struct {
	r           *Repository
	args        []string
	env         []string
	stdIn       io.Reader
	unsetGitDir bool
}

// executor initializes a new executor instance to run a Git command with the
// specified arguments.
func (r *Repository) executor(args ...string) *executor { _ = "STUB: not implemented"; return nil }

// withEnv adds the specified environment variables. Each environment variable
// must be specified in the form of `key=value`.
func (e *executor) withEnv(env ...string) *executor { _ = "STUB: not implemented"; return nil }

// withoutGitDir ensures the executor doesn't auto-set the --git-dir flag to the
// executed command.
func (e *executor) withoutGitDir() *executor { _ = "STUB: not implemented"; return nil }

// withStdIn sets the contents of stdin to be passed in to the command.
func (e *executor) withStdIn(stdIn *bytes.Buffer) *executor { _ = "STUB: not implemented"; return nil }

// executeString runs the constructed Git command and returns the contents of
// stdout.  Leading and trailing spaces and newlines are removed. This function
// should be used almost every time; the only exception is when the output is
// desired without any processing such as the removal of space characters.
func (e *executor) executeString() (string, error) { _ = "STUB: not implemented"; return "", nil }

// execute runs the constructed Git command and returns the raw stdout and
// stderr contents. It adds the `--git-dir` argument if the repository has a
// path set.
func (e *executor) execute() (io.Reader, io.Reader, error) {
	_ = "STUB: not implemented"
	return *new(io.Reader), *new(io.Reader), nil
}

//nolint:gosec

// force git to the C (and thus english) locale
