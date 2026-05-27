// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package display

import (
	"io"

	"github.com/gittuf/gittuf/internal/common/set"
	"github.com/gittuf/gittuf/internal/rsl"
	"github.com/gittuf/gittuf/pkg/gitinterface"
)

type options struct {
	refs *set.Set[string]
}

type Option func(*options)

func WithReferences(refs []string) Option { _ = "STUB: not implemented"; return *new(Option) }

// RSLLog implements the display function for `gittuf rsl log`.
func RSLLog(repo *gitinterface.Repository, writer io.WriteCloser, opts ...Option) error {
	_ = "STUB: not implemented"
	//nolint:errcheck
	return nil
}

// assume an entry has a parent

// only reachable when err is ErrRSLEntryNotFound
// Now we know the iteratorEntry does not have a parent

// Skip this entry if it's not for the specified ref. Note that
// we still want to track annotation entries for this entry (if
// there are any) since they may apply to other entries that we
// do want to display.

// We return nil here to avoid noisy output when the writer is
// unexpectedly closed, such as by killing the pager

// Skip this entry if it's not for the specified ref. Note that
// we still want to track annotation entries for this entry (if
// there are any) since they may apply to other entries that we
// do want to display.

// We return nil here to avoid noisy output when
// the writer is unexpectedly closed, such as by
// killing the pager

// We're done

// writeRSLReferenceEntry prepares the output for the given entry and its
// annotations. It then writes the output to the provided writer. If hasParent
// is false, then the prepared output for the entry has a single trailing
// newline. Otherwise, an additional newline is added to separate entries from
// one another.
func writeRSLReferenceEntry(writer io.WriteCloser, entry *rsl.ReferenceEntry, annotations []*rsl.AnnotationEntry, hasParent bool) error {
	_ = "STUB: not implemented"
	/* Output format:
	   entry <entryID> (skipped)

	     Ref:    <refName>
	     Target: <targetID>
	     Number: <number>

	       Annotation ID: <annotationID>
	       Skip:          <yes/no>
	       Number:        <number>
	       Message:
	         <message>

	       Annotation ID: <annotationID>
	       Skip:          <yes/no>
	       Number:        <number>
	       Message:
	         <message>
	*/return nil
}

// single trailing newline by default

// extra newline for all intermediate (i.e., not last) entries

func writeRSLPropagationEntry(writer io.WriteCloser, entry *rsl.PropagationEntry, hasParent bool) error {
	_ = "STUB: not implemented"
	/* Output format:
	   propagation entry <entryID>
	     Ref:           <refName>
	     Target:        <targetID>
		 UpstreamRepo:  <upstreamRepoLocation>
		 UpstreamEntry: <upstreamEntryID>
	     Number:        <number>
	*/return nil
}

// single trailing newline by default

// extra newline for all intermediate (i.e., not last) entries
