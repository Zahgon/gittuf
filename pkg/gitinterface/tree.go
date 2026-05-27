// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package gitinterface

import (
	"errors"
	"os"
)

var (
	ErrTreeDoesNotHavePath             = errors.New("tree does not have requested path")
	ErrCopyingBlobIDsDoNotMatch        = errors.New("blob ID in local repository does not match upstream repository")
	ErrCannotCreateSubtreeIntoRootTree = errors.New("subtree path target cannot be empty or root of tree")
)

func (r *Repository) EmptyTree() (Hash, error) { _ = "STUB: not implemented"; return *new(Hash), nil }

// GetPathIDInTree returns the Git ID pointed to by the path in the specified
// tree if the path exists. If not, a corresponding error is returned.  For
// example, if the tree contains a single blob `foo/bar/baz`, querying the ID
// for `foo/bar/baz` will return the blob ID for baz. Querying the ID for
// `foo/bar` will return the intermediate tree ID for bar, while querying for
// `foo/baz` will return an error.
func (r *Repository) GetPathIDInTree(treePath string, treeID Hash) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// GetTreeItems returns the items in a specified Git tree without recursively
// expanding subtrees.
func (r *Repository) GetTreeItems(treeID Hash) (map[string]Hash, error) {
	_ = "STUB: not implemented"
	// From Git 2.36, we can use --format here. However, it appears a not
	// insignificant number of developers are still on Git 2.34.1, a side effect
	// of being on Ubuntu 22.04. 22.04 is still widely used in WSL2 environments.
	// So, we're removing --format and parsing the output differently to handle
	// the extra information for each entry we don't need.
	return nil, nil
}

// alternatively, just check if treeID is empty tree?

// Without --format, the output is in the following format:
// <mode> SP <type> SP <object> TAB <file>
// From: https://git-scm.com/docs/git-ls-tree/2.34.1#_output_format

// entrySplit[0] is <mode> -- discard
// entrySplit[1] is <type> -- discard
// entrySplit[2] is <object> TAB <file> -- keep

// <object> is really the object ID

// GetAllFilesInTree returns all filepaths and the corresponding blob hashes in
// the specified tree.
func (r *Repository) GetAllFilesInTree(treeID Hash) (map[string]Hash, error) {
	_ = "STUB: not implemented"
	// From Git 2.36, we can use --format here. However, it appears a not
	// insignificant number of developers are still on Git 2.34.1, a side effect
	// of being on Ubuntu 22.04. 22.04 is still widely used in WSL2 environments.
	// So, we're removing --format and parsing the output differently to handle
	// the extra information for each entry we don't need.
	return nil, nil
}

// alternatively, just check if treeID is empty tree?

// Without --format, the output is in the following format:
// <mode> SP <type> SP <object> TAB <file>
// From: https://git-scm.com/docs/git-ls-tree/2.34.1#_output_format

// entrySplit[0] is <mode> -- discard
// entrySplit[1] is <type> -- discard
// entrySplit[2] is <object> TAB <file> -- keep

// <object> is really the object ID

// GetMergeTree computes the merge tree for the commits passed in. The tree is
// not written to the object store. Assuming a typical merge workflow, the first
// commit is expected to be the tip of the base branch. As such, the second
// commit is expected to be merged into the first. If the first commit is zero,
// the second commit's tree is returned.
func (r *Repository) GetMergeTree(commitAID, commitBID Hash) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// fast-forward merge -> use tree ID from commitB

// Only commitB needs to be non-zero, we can allow fast-forward merges when
// the base commit is zero. So, check this only after above

// Older Git versions do not support merge-tree, and, as such, require
// quite a long workaround to find what the merge tree is. This
// workaround boils down to:
// Create new branch > Merge into said branch > Extract tree hash

// Attempt to abort the merge in all cases as a failsafe

// Switch back to the branch the user was on

// CreateSubtreeFromUpstreamRepository accepts an upstream repository handler
// and a commit ID in the upstream repository. This information is used to copy
// the entire contents of the commit's Git tree into the specified localPath in
// the localRef. A new commit is added to localRef with the changes made to
// localPath. localPath represents a directory path where the changes are copied
// to. Existing items in that directory are overwritten in the subsequently
// created commit in localRef. localPath must be specified, if left blank (say
// to imply copying into the root directory of the downstream repository),
// creating a subtree will fail.
func (r *Repository) CreateSubtreeFromUpstreamRepository(upstream *Repository, upstreamCommitID Hash, upstreamPath, localRef, localPath string) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// Ignore entries for `localPath` to account for upstream deletions
// If localPath is foo/, we want to ignore all items under foo/
// If localPath is foo, we want to ignore all items under foo/
// If localPath is foo, we DO NOT want to remove all items under foobar/
// So, add the / suffix if necessary to localPath

// Create list of TreeEntry objects representing all blobs except those
// currently under localPath

// Remove trailing "/" now

// If upstreamPath is empty, then the entire tree is copied over,
// otherwise, identify the subtree to copy over

// Use existing intermediate tree

// We have to create the intermediate tree for localPath

// if blob already exists, we don't need to carry out expensive
// read/write

// add blob to entries, with the path including the localPath prefix

// TODO: this doesn't support detached git dir

//nolint:errcheck

// TreeBuilder is used to create multi-level trees in a repository.  Based on
// `buildTreeHelper` in go-git.
type TreeBuilder struct {
	repo    *Repository
	trees   map[string]*entryTree
	entries map[string]TreeEntry
}

func NewTreeBuilder(repo *Repository) *TreeBuilder { _ = "STUB: not implemented"; return nil }

// WriteTreeFromEntries accepts list of TreeEntry representations, and returns
// the Git ID of the tree that contains these entries. It constructs the
// required intermediate trees.
func (t *TreeBuilder) WriteTreeFromEntries(files []TreeEntry) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// identifyIntermediates identifies the intermediate trees that must be
// constructed for the specified path.
func (t *TreeBuilder) identifyIntermediates(entry TreeEntry) { _ = "STUB: not implemented"; return }

// populateTree populates tree and entry information for each tree that must be
// created.
func (t *TreeBuilder) populateTree(parent, fullPath string, entry TreeEntry) {
	_ = "STUB: not implemented"
	return
}

// => This is a leaf node
// However, gitID _may_ be a tree ID, and we've inserted an existing
// tree object as a subtree here, we want to support this so that we
// don't have to recreate trees that already exist

// gitID represents tree

// gitID is not for a tree

// => This is an intermediate node, has to be a tree that we must build

// writeTrees recursively stores each tree that must be created in the
// repository's object store. It returns the ID of the tree created at each
// invocation.
func (t *TreeBuilder) writeTrees(parent string, tree *entryTree) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// The tree already exists and we don't need to write it again.

// writeTree creates a tree in the repository for the specified entries. It
// only supports a typical blob with permission 0o644 and a subtree. This is
// because it is only intended for use with gittuf specific metadata and tests.
// Generic tree creation is left to invocations of the Git binary by the user.
func (t *TreeBuilder) writeTree(entries []TreeEntry) (Hash, error) {
	_ = "STUB: not implemented"
	return *new(Hash), nil
}

// this is very opinionated about the modes right now because the plan
// is to use it for gittuf metadata, which requires regular files and
// subdirectories

// TODO: support entryBlob's permissions here

// TreeEntry represents an entry in a Git tree.
type TreeEntry interface {
	getName() string
	getID() Hash
}

// entryTree implements TreeEntry and indicates the entry is for a Git tree.
type entryTree struct {
	name          string
	gitID         Hash
	alreadyExists bool
	entries       []TreeEntry
}

func (e *entryTree) getName() string { _ = "STUB: not implemented"; return "" }

func (e *entryTree) getID() Hash {
	_ = "STUB: not implemented"

	// NewEntryTree creates a TreeEntry that represents a Git tree. If the tree
	// doesn't exist, i.e., it must be created, gitID must be set to ZeroHash. The
	// name must be set to the full path of the tree object.
	return *new(Hash)
}

func NewEntryTree(name string, gitID Hash) TreeEntry {
	_ = "STUB: not implemented"
	return *new(TreeEntry)
}

// entryBlob implements TreeEntry and indicates the entry is for a Git blob.
type entryBlob struct {
	name        string
	gitID       Hash
	permissions os.FileMode //nolint:unused
}

func (e *entryBlob) getName() string { _ = "STUB: not implemented"; return "" }

func (e *entryBlob) getID() Hash {
	_ = "STUB: not implemented"

	// NewEntryBlob creates a TreeEntry that represents a Git blob.
	return *new(Hash)
}

func NewEntryBlob(name string, gitID Hash) TreeEntry {
	_ = "STUB: not implemented"
	return *new(TreeEntry)
}

// NewEntryBlobWithPermissions creates a TreeEntry that represents a Git blob.
// The permissions parameter can be used to set custom permissions.
func NewEntryBlobWithPermissions(name string, gitID Hash, permissions os.FileMode) TreeEntry {
	_ = "STUB: not implemented"
	return *new(TreeEntry)
}

// ensureIsTree is a helper to check that the ID represents a Git tree
// object.
func (r *Repository) ensureIsTree(treeID Hash) error { _ = "STUB: not implemented"; return nil }
