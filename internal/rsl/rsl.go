// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package rsl

import (
	"errors"

	"github.com/gittuf/gittuf/internal/tuf"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

const (
	Ref       = "refs/gittuf/reference-state-log"
	NumberKey = "number"

	ReferenceEntryHeader = "RSL Reference Entry"
	RefKey               = "ref"
	TargetIDKey          = "targetID"

	AnnotationEntryHeader      = "RSL Annotation Entry"
	AnnotationMessageBlockType = "MESSAGE"
	BeginMessage               = "-----BEGIN MESSAGE-----"
	EndMessage                 = "-----END MESSAGE-----"
	EntryIDKey                 = "entryID"
	SkipKey                    = "skip"

	PropagationEntryHeader = "RSL Propagation Entry"
	UpstreamRepositoryKey  = "upstreamRepository"
	UpstreamEntryIDKey     = "upstreamEntryID"

	remoteTrackerRef       = "refs/remotes/%s/gittuf/reference-state-log"
	gittufNamespacePrefix  = "refs/gittuf/"
	gittufPolicyStagingRef = "refs/gittuf/policy-staging"
)

var (
	ErrRSLEntryNotFound                             = errors.New("unable to find RSL entry")
	ErrRSLBranchDetected                            = errors.New("potential RSL branch detected, entry has more than one parent")
	ErrInvalidRSLEntry                              = errors.New("RSL entry has invalid format or is of unexpected type")
	ErrRSLEntryDoesNotMatchRef                      = errors.New("RSL entry does not match requested ref")
	ErrNoRecordOfCommit                             = errors.New("commit has not been encountered before")
	ErrInvalidGetLatestReferenceUpdaterEntryOptions = errors.New("invalid options presented for getting latest reference updater entry (are both before or until conditions set or is the before number less than the until number?)")
	ErrCannotUseEntryNumberFilter                   = errors.New("current RSL entries are not numbered, cannot use number range options")
	ErrInvalidUntilEntryNumberCondition             = errors.New("cannot meet until entry number condition")
)

// RemoteTrackerRef returns the remote tracking ref for the specified remote
// name. For example, for 'origin', the remote tracker ref is
// 'refs/remotes/origin/gittuf/reference-state-log'.
func RemoteTrackerRef(remote string) string { _ = "STUB: not implemented"; return "" }

// Entry is the abstract representation of an object in the RSL.
type Entry interface {
	GetID() gitinterface.Hash
	Commit(*gitinterface.Repository, bool) error
	GetNumber() uint64
	createCommitMessage(bool) (string, error)
}

// ReferenceUpdaterEntry represents RSL entry types that can record an update to
// a Git reference. Some examples are the reference entry and the propagation
// entry.
type ReferenceUpdaterEntry interface {
	Entry
	GetRefName() string
	GetTargetID() gitinterface.Hash
}

// ReferenceEntry represents a record of a reference state in the RSL. It
// implements the Entry interface.
type ReferenceEntry struct {
	// ID contains the Git hash for the commit corresponding to the entry.
	ID gitinterface.Hash

	// RefName contains the Git reference the entry is for.
	RefName string

	// TargetID contains the Git hash for the object expected at RefName.
	TargetID gitinterface.Hash

	// Number contains a strictly increasing number that hints at entry ordering.
	Number uint64
}

// NewReferenceEntry returns a ReferenceEntry object for a normal RSL entry.
func NewReferenceEntry(refName string, targetID gitinterface.Hash) *ReferenceEntry {
	_ = "STUB: not implemented"
	return nil
}

func (e *ReferenceEntry) GetID() gitinterface.Hash {
	_ = "STUB: not implemented"
	return *new(gitinterface.Hash)
}

func (e *ReferenceEntry) GetRefName() string { _ = "STUB: not implemented"; return "" }

func (e *ReferenceEntry) GetTargetID() gitinterface.Hash {
	_ = "STUB: not implemented"

	// Commit creates a commit object in the RSL for the ReferenceEntry. The
	// function looks up the latest committed entry in the RSL and increments the
	// number in the new entry. If a parent entry does not exist or the parent
	// entry's number is 0 (unset), the current entry's number is set to 1. The
	// numbering starts from 1 as 0 is used to signal the lack of numbering.
	return *new(gitinterface.Hash)
}

func (e *ReferenceEntry) Commit(repo *gitinterface.Repository, sign bool) error {
	_ = "STUB: not implemented"
	return nil
}

// we have an error return for annotations, always nil here

// CommitUsingSpecificKey creates a commit object in the RSL for the
// ReferenceEntry. The commit is signed using the provided PEM encoded SSH or
// GPG private key. This is only intended for use in gittuf's developer mode or
// in tests. The function looks up the latest committed entry in the RSL and
// increments the number in the new entry. If a parent entry does not exist or
// the parent entry's number is 0 (unset), the current entry's number is set to
// 1. The numbering starts from 1 as 0 is used to signal the lack of numbering.
func (e *ReferenceEntry) CommitUsingSpecificKey(repo *gitinterface.Repository, signingKeyBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// we have an error return for annotations, always nil here

func (e *ReferenceEntry) GetNumber() uint64 {
	_ = "STUB: not implemented"

	// Skipped returns true if any of the annotations mark the entry as
	// to-be-skipped.
	return 0
}

func (e *ReferenceEntry) SkippedBy(annotations []*AnnotationEntry) bool {
	_ = "STUB: not implemented"
	return false
}

func (e *ReferenceEntry) setEntryNumber(repo *gitinterface.Repository) error {
	_ = "STUB: not implemented"
	return nil
}

// First entry

func (e *ReferenceEntry) createCommitMessage(includeNumber bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// commitWithoutNumber is used to test the RSL's support for entry numbers in
// repositories that switch from not having numbered entries to having numbered
// entries.
func (e *ReferenceEntry) commitWithoutNumber(repo *gitinterface.Repository) error {
	_ = "STUB: not implemented"
	return nil
}

// we have an error return for annotations, always nil here

// AnnotationEntry is a type of RSL record that references prior items in the
// RSL. It can be used to add extra information for the referenced items.
// Annotations can also be used to "skip", i.e. revoke, the referenced items. It
// implements the Entry interface.
type AnnotationEntry struct {
	// ID contains the Git hash for the commit corresponding to the annotation.
	ID gitinterface.Hash

	// RSLEntryIDs contains one or more Git hashes for the RSL entries the annotation applies to.
	RSLEntryIDs []gitinterface.Hash

	// Skip indicates if the RSLEntryIDs must be skipped during gittuf workflows.
	Skip bool

	// Message contains any messages or notes added by a user for the annotation.
	Message string

	// Number contains a strictly increasing number that hints at entry ordering.
	Number uint64
}

// NewAnnotationEntry returns an Annotation object that applies to one or more
// prior RSL entries.
func NewAnnotationEntry(rslEntryIDs []gitinterface.Hash, skip bool, message string) *AnnotationEntry {
	_ = "STUB: not implemented"
	return nil
}

func (a *AnnotationEntry) GetID() gitinterface.Hash {
	_ = "STUB: not implemented"

	// Commit creates a commit object in the RSL for the Annotation. The function
	// looks up the latest committed entry in the RSL and increments the number in
	// the new entry. If a parent entry does not exist or the parent entry's number
	// is 0 (unset), the current entry's number is set to 1. The numbering starts
	// from 1 as 0 is used to signal the lack of numbering.
	return *new(gitinterface.Hash)
}

func (a *AnnotationEntry) Commit(repo *gitinterface.Repository, sign bool) error {
	_ = "STUB: not implemented"
	// Check if referred entries exist in the RSL namespace.
	return nil
}

// CommitUsingSpecificKey creates a commit object in the RSL for the
// AnnotationEntry. The commit is signed using the provided PEM encoded SSH or
// GPG private key. This is only intended for use in gittuf's developer mode or
// in tests. The function looks up the latest committed entry in the RSL and
// increments the number in the new entry. If a parent entry does not exist or
// the parent entry's number is 0 (unset), the current entry's number is set to
// 1. The numbering starts from 1 as 0 is used to signal the lack of numbering.
func (a *AnnotationEntry) CommitUsingSpecificKey(repo *gitinterface.Repository, signingKeyBytes []byte) error {
	_ = "STUB: not implemented"
	// Check if referred entries exist in the RSL namespace.
	return nil
}

func (a *AnnotationEntry) GetNumber() uint64 {
	_ = "STUB: not implemented"

	// RefersTo returns true if the specified entryID is referred to by the
	// annotation.
	return 0
}

func (a *AnnotationEntry) RefersTo(entryID gitinterface.Hash) bool {
	_ = "STUB: not implemented"
	return false
}

func (a *AnnotationEntry) setEntryNumber(repo *gitinterface.Repository) error {
	_ = "STUB: not implemented"
	return nil
}

// First entry -> can an annotation actually be first? TODO

func (a *AnnotationEntry) createCommitMessage(includeNumber bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// commitWithoutNumber is used to test the RSL's support for entry numbers in
// repositories that switch from not having numbered entries to having numbered
// entries.
func (a *AnnotationEntry) commitWithoutNumber(repo *gitinterface.Repository) error {
	_ = "STUB: not implemented"
	// Check if referred entries exist in the RSL namespace.
	return nil
}

// PropagationEntry represents a record of execution of gittuf's repository
// propagation workflow. It indicates which reference was updated with an
// upstream repository's contents, as well as details about the upstream
// repository such as its location and the specific entry whose contents were
// propagated.
type PropagationEntry struct {
	// ID contains the Git hash for the commit corresponding to the entry.
	ID gitinterface.Hash

	// RefName contains the Git reference the entry is for.
	RefName string

	// TargetID contains the Git hash for the object expected at RefName.
	TargetID gitinterface.Hash

	// UpstreamRepository records the location of the upstream repository.
	UpstreamRepository string

	// UpstreamEntryID records the upstream repository's RSL entry ID whose
	// contents were propagated.
	UpstreamEntryID gitinterface.Hash

	// Number contains a strictly increasing number that hints at entry ordering.
	Number uint64
}

func NewPropagationEntry(refName string, targetID gitinterface.Hash, upstreamRepository string, upstreamEntryID gitinterface.Hash) *PropagationEntry {
	_ = "STUB: not implemented"
	return nil
}

func (e *PropagationEntry) GetID() gitinterface.Hash {
	_ = "STUB: not implemented"
	return *new(gitinterface.Hash)
}

func (e *PropagationEntry) GetRefName() string { _ = "STUB: not implemented"; return "" }

func (e *PropagationEntry) GetTargetID() gitinterface.Hash {
	_ = "STUB: not implemented"

	// Commit creates a commit object in the RSL for the PropagationEntry. The
	// function looks up the latest committed entry in the RSL and increments the
	// number in the new entry. If a parent entry does not exist or the parent
	// entry's number is 0 (unset), the current entry's number is set to 1. The
	// numbering starts from 1 as 0 is used to signal the lack of numbering.
	return *new(gitinterface.Hash)
}

func (e *PropagationEntry) Commit(repo *gitinterface.Repository, sign bool) error {
	_ = "STUB: not implemented"
	return nil
}

// we have an error return for annotations, always nil here

// CommitUsingSpecificKey creates a commit object in the RSL for the
// PropagationEntry. The commit is signed using the provided PEM encoded SSH or
// GPG private key. This is only intended for use in gittuf's developer mode or
// in tests. The function looks up the latest committed entry in the RSL and
// increments the number in the new entry. If a parent entry does not exist or
// the parent entry's number is 0 (unset), the current entry's number is set to
// 1. The numbering starts from 1 as 0 is used to signal the lack of numbering.
func (e *PropagationEntry) CommitUsingSpecificKey(repo *gitinterface.Repository, signingKeyBytes []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// we have an error return for annotations, always nil here

func (e PropagationEntry) GetNumber() uint64 { _ = "STUB: not implemented"; return 0 }

func (e *PropagationEntry) setEntryNumber(repo *gitinterface.Repository) error {
	_ = "STUB: not implemented"
	return nil
}

// First entry

func (e *PropagationEntry) createCommitMessage(includeNumber bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetEntry returns the entry corresponding to entryID.
func GetEntry(repo *gitinterface.Repository, entryID gitinterface.Hash) (Entry, error) {
	_ = "STUB: not implemented"
	return *new(Entry), nil
}

// GetParentForEntry returns the entry's parent RSL entry.
func GetParentForEntry(repo *gitinterface.Repository, entry Entry) (Entry, error) {
	_ = "STUB: not implemented"
	return *new(Entry), nil
}

// We don't need to check the parent's Number here because it was
// checked when this was set in the cache

// parent entry has to be 0

// parent entry has to be 1 less than entry

// GetNonGittufParentReferenceUpdaterEntryForEntry returns the first RSL
// reference updater entry starting from the specified entry's parent that is
// not for the gittuf namespace.
func GetNonGittufParentReferenceUpdaterEntryForEntry(repo *gitinterface.Repository, entry Entry) (ReferenceUpdaterEntry, []*AnnotationEntry, error) {
	_ = "STUB: not implemented"
	return *new(ReferenceUpdaterEntry), nil, nil
}

// we've found the target entry, stop walking the RSL

// GetLatestEntry returns the latest entry available locally in the RSL.
func GetLatestEntry(repo *gitinterface.Repository) (Entry, error) {
	_ = "STUB: not implemented"
	return *new(Entry), nil
}

// GetLatestReferenceUpdaterEntry returns the latest reference updater entry in
// the local RSL that matches the specified conditions.
func GetLatestReferenceUpdaterEntry(repo *gitinterface.Repository, opts ...GetLatestReferenceUpdaterEntryOption) (ReferenceUpdaterEntry, []*AnnotationEntry, error) {
	_ = "STUB: not implemented"
	return *new(ReferenceUpdaterEntry), nil, nil
}

// Only one of the Before options can be set

// Only one of the Until options can be set

// Sanity check before / until number conditions

// The repository doesn't use numbers yet

// Do initial walk if either before condition is set

// we've found the before anchor entry, track it if it's an
// annotation

// Set it to parent as this is the first entry considered below
// While this entry may match equal until condition, that's fine
// as the until condition is inclusive

// Only reference entry can be skipped

// SkippedBy ensures only the applicable
// annotations that refer to the entry
// are used

// We've found the target entry, stop walking the RSL

// GetFirstEntry returns the very first entry in the RSL. It is expected to be a
// reference updater entry as the first entry in the RSL cannot be an
// annotation.
func GetFirstEntry(repo *gitinterface.Repository) (ReferenceUpdaterEntry, []*AnnotationEntry, error) {
	_ = "STUB: not implemented"
	return *new(ReferenceUpdaterEntry), nil, nil
}

// GetFirstReferenceEntryForRef returns the very first entry in the RSL for the
// specified ref. It is expected to be a reference entry as the first entry in
// the RSL for a reference cannot be an annotation.
func GetFirstReferenceUpdaterEntryForRef(repo *gitinterface.Repository, targetRef string) (ReferenceUpdaterEntry, []*AnnotationEntry, error) {
	_ = "STUB: not implemented"
	return *new(ReferenceUpdaterEntry), nil, nil
}

// SkipAllInvalidReferenceEntriesForRef identifies invalid RSL reference entries.
// Each invalid entry points to a target that is not reachable for the current
// target of the same reference, indicating that history has been rewritten via a
// rebase for the reference. After the invalid entries are identified, an annotation
// entry is created that marks all of these entries as to be skipped.
func SkipAllInvalidReferenceEntriesForRef(repo *gitinterface.Repository, targetRef string, signCommit bool) error {
	_ = "STUB: not implemented"
	return nil
}

// We don't have a parent to check if invalid
// So we assume the current one is valid
// TODO: should we cross reference state of the branch?

// GetFirstReferenceUpdaterEntryForCommit returns the first reference entry in
// the RSL that either records the commit itself or a descendent of the commit.
// This establishes the first time a commit was seen in the repository,
// irrespective of the ref it was associated with, and we can infer things like
// the active developers who could have signed the commit.
func GetFirstReferenceUpdaterEntryForCommit(repo *gitinterface.Repository, commitID gitinterface.Hash) (ReferenceUpdaterEntry, []*AnnotationEntry, error) {
	_ = "STUB: not implemented"
	// We check entries in pairs. In the initial case, we have the latest entry
	// and its parent. At all times, the parent in the pair is being tested.
	// If the latest entry is a descendant of the target commit, we start
	// checking the parent. The first pair where the parent entry is not
	// descended from the target commit, we return the other entry in the pair.
	return *new(ReferenceUpdaterEntry), nil, nil
}

// GetReferenceUpdaterEntriesInRange returns a list of reference entries between
// the specified range and a map of annotations that refer to each reference
// entry in the range. The annotations map is keyed by the ID of the reference
// entry, with the value being a list of annotations that apply to that
// reference entry.
func GetReferenceUpdaterEntriesInRange(repo *gitinterface.Repository, firstID, lastID gitinterface.Hash) ([]ReferenceUpdaterEntry, map[string][]*AnnotationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// GetReferenceUpdaterEntriesInRangeForRef returns a list of reference entries
// for the ref between the specified range and a map of annotations that refer
// to each reference entry in the range. The annotations map is keyed by the ID
// of the reference entry, with the value being a list of annotations that apply
// to that reference entry.
func GetReferenceUpdaterEntriesInRangeForRef(repo *gitinterface.Repository, firstID, lastID gitinterface.Hash, refName string) ([]ReferenceUpdaterEntry, map[string][]*AnnotationEntry, error) {
	_ = "STUB: not implemented"
	// We have to iterate from latest to get the annotations that refer to the
	// last requested entry
	return nil, nil, nil
}

// Until we find the entry corresponding to lastID, we just store
// annotations

// Here, all items are relevant until the one corresponding to first is
// found

// It's a relevant entry if:
// a) there's no refName set, or
// b) the entry's refName matches the set refName, or
// c) the entry is for a gittuf namespace

// Handle the item corresponding to first explicitly
// If it's an annotation, ignore it as it refers to something before the
// range we care about

// It's a relevant entry if:
// a) there's no refName set, or
// b) the entry's refName matches the set refName, or
// c) the entry is for a gittuf namespace

// For each annotation, add the entry to each relevant entry it refers to
// Process annotations in reverse order so that annotations are listed in
// order of occurrence in the map

// Annotation is relevant because the entry it refers to was in
// the specified range

// Reverse entryStack so that it's in order of occurrence rather than in
// order of walking back the RSL

// PropagateChangesFromUpstreamRepository executes gittuf's propagation workflow
// to create a subtree of the contents of an upstream repository's reference
// into the specified reference and path in the downstream repository.
func PropagateChangesFromUpstreamRepository(downstreamRepo, upstreamRepo *gitinterface.Repository, details []tuf.PropagationDirective, sign bool) error {
	_ = "STUB: not implemented"
	// FIXME: We assume here that downstreamRepo and upstreamRepo have their
	// gittuf refs already synced.
	return nil
}

// We want to check if propagation is necessary
// What if it's already been propagated?

// TODO: handle divergence from latest RSL entry for ref downstream?

// TODO: should we handle this differently?

// TODO: should we handle this differently?

// Nothing to do

// TODO: error management should revert propagation entries?
// atomicity?

func parseRSLEntryText(id gitinterface.Hash, text string) (Entry, error) {
	_ = "STUB: not implemented"
	return *new(Entry), nil
}

func parseReferenceEntryText(id gitinterface.Hash, text string) (*ReferenceEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseAnnotationEntryText(id gitinterface.Hash, text string) (*AnnotationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// rest doesn't seem to work when the PEM block is at the end of text, see: https://go.dev/play/p/oZysAfemA-v

func parsePropagationEntryText(id gitinterface.Hash, text string) (*PropagationEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The location may also have `:`, so we need to handle all items in ls

func filterAnnotationsForRelevantAnnotations(allAnnotations []*AnnotationEntry, entryID gitinterface.Hash) []*AnnotationEntry {
	_ = "STUB: not implemented"
	return nil
}

func isRelevantGittufRef(refName string) bool { _ = "STUB: not implemented"; return false }
