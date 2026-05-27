// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"

	"github.com/gittuf/gittuf/experimental/gittuf"
)

// handleSSH implements the helper for remotes configured to use SSH. For this
// transport, we invoke the installed ssh binary to interact with the remote.
func handleSSH(ctx context.Context, repo *gittuf.Repository, remoteName, url string) (map[string]string, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

// 0 is the connection [user@]host, 1 is the repo

// Scan git-remote-gittuf stdin for commands from the parent process

/*
	For SSH, we have several options wrt capabilities. First, we
	could just implement fetch and push. These are v0/v1 protocols.
	The issue here is that while push is fine, fetch effectively
	fetches _all_ refs it sees on the remote via list. Additionally,
	using v2 protocol where possible seems good for efficiency
	improvements hinted at by the docs.

	The connect capability sets up a bidirectional connection with
	the server. It can handle both fetches and pushes; depending on
	what's happening, either upload-pack or receive-pack must be
	invoked on the server. This is fine for fetch operations.
	However, for push, we can tell the server to set
	refs/gittuf/<whatever> to the object. However, we do not control
	the invocation of git pack-objects --stdout. Git (which invokes
	us) invokes pack-objects separately, and routes its stdout into
	the transport's stdin to transmit the packfile bytes.

	In summary, we cannot use a combination of fetch and push, and
	we cannot use connect. What about stateless-connect?  This is
	part of the v2 protocol and can only handle fetches at the
	moment. It's marked as experimental, which is something we want
	to be wary about with new Git versions.  There may well be
	breaking changes here, given that the only intended user of this
	command is other Git tooling.

	stateless-connect is quite easy to work with to handle the fetch
	aspects. In addition, we implement the push capability. Here,
	Git tells us the refspecs that must be pushed. We are separately
	responsible for actually sending the packfile(s). So, the
	solution is that we create RSL entries for each requested ref,
	and include the gittuf objects in the packfile. Thus, we specify
	stateless-connect and push as the two capabilities supported by
	this helper.
*/

/*
	When we see stateless-connect, right now we know this means
	a fetch is underway.

	us: ssh -o SendEnv=GIT_PROTOCOL <url> 'git-upload-pack <repo>'
	ssh:
		if v2 {
			server capabilities
		} else {
			server capabilities
			refs and their states
		}
	Assuming v2:
	us (to ssh): ls-refs // add gittuf prefix
	ssh: refs and their states
	us (to git): output of ls-refs
	git: fetch, wants, haves
	us (to ssh): fetch, wants, haves // add gittuf wants
	ssh: acks (optionally triggers another round of wants, haves)
	ssh: packfile
	us (to git): acks, packfile

	Assuming v0/v1:
	git: wants, haves // NO FETCH HERE IIRC
	us (to ssh): wants, haves // add gittuf wants
	ssh: acks, packfile
	us (to git): acks, packfile

	Notes:
		* v0/v1 of the pack protocol is only partially supported
		  here.
		* Once the service is invoked, all messages are wrapped
		  in the packet-line format.
		* In the v0/v1 format, each line is packet encoded, and
		  the entire message is in turn packet encoded for
		  wants/haves.
		* The flushPkt is commonly used to signify end of a
		  message.
		* The endOfReadPkt is sent at the end of the packfile
		  transmission.
*/

// This allows us to request GIT_PROTOCOL v2

// with stateless-connect, it's only fetches

// Crafting ssh subprocess for fetches
//nolint:gosec
// Add env var for GIT_PROTOCOL v2

// We want to inspect the helper's stdout for gittuf ref statuses

// We want to interpose with the helper's stdin by passing in
// extra refs etc.

// Indicate connection established successfully

// Read from remote service
// TODO: we may need nested infinite loops here

// TODO: handle git protocol v0/v1
// If server doesn't support v2, as soon as we connect,
// it tells us the ref statuses

// check for end of message

// In protocol v2, this should now go to our parent process
// requesting ls-refs

// Add ref-prefix refs/gittuf/ to the ls-refs command before
// flush, but only if the client sent ref-prefixes

// Check for end of message

// In the curl transport, we also look out for endOfReadPkt
// However, this has been a bit flakey
// So when we see the flushPkt, we'll also write the
// endOfReadPkt ourselves

// remove pkt length prefix

// If the gittuf ref is the very first, then there will be
// additional information in the output after a null byte.
// However, this is unlikely as HEAD is typically the first.

// drop everything from null byte onwards

// Write output to parent process

// For a stateless connection, we must
// also add the endOfRead packet
// ourselves

// At this point, we enter the haves / wants negotiation, which is
// followed usually by the remote sending a packfile with the
// requested Git objects.

// We add the gittuf specific objects as wants. We don't have to
// specify haves as Git automatically specifies all the objects it
// has regardless of what refs they're reachable via.

// Read in command from parent process -> this should be
// command=fetch with protocol v2

// track this in case there are multiple rounds of negotiation

// We're done but we need to exit gracefully

// We only write gittuf specific haves and wants when we
// haven't already written them. We track this because
// in multiple rounds of negotiations, we only want to
// write them the first time.

// indicate we
// want the
// gittuf obj

// indicate we
// have the
// gittuf obj

// Take out this ref as
// something for us to
// update or add wants
// for

// Read from remote if wants are done
// We may need to scan multiple times for inputs, which is why
// this flag is used

// TODO: find something cleaner to terminate

// Send along to parent process

// we see this at the end

// Go back for more input

/*
	git: list for-push // wants to know remote ref statuses
	us: ssh...git-receive-pack
	ssh: list of refs
	us (to git): list of refs // trailing newline

	git: push cmds
	us: track list of push cmds, create RSL entry for each
	us (to ssh): push cmds (receive-pack format) // also track oldTip for eachRef
	us (to ssh): git pack-objects > ssh // include object range desired
*/

// This allows us to request GIT_PROTOCOL v2

// with stateless-connect, it's only fetches

// Crafting ssh subprocess for pushes
//nolint:gosec
// Add env var for GIT_PROTOCOL v2

// We want to inspect the helper's stdout for gittuf ref statuses

// We want to interpose with the helper's stdin by passing in
// extra refs etc.

// TODO: does this need a nested loop?

// TODO: do we need endOfReadPkt check?

// remove length prefix

// remove config string passed after null byte

// We don't use ref, instead we use refAdSplit[1]. This
// allows us to propagate remote capabilities to the parent
// process
//nolint:gosec

// Add trailing new line as we're bridging git-receive-pack
// output with git remote helper output

// We explicitly push the RSL ref below
// because we need to know what its tip
// will be after all other refs are
// pushed.

// TODO: skipping propagation; invoke it once total instead of per ref

// report-status-v2 indicates we want the result for each pushed ref
// atomic indicates either all must be successful or none
// object-format indicates SHA-1 vs SHA-256 repo
// agent indicates the version of the local git client (most of the time)
// Note: we explicitly don't use the sideband here
// because of inconsistencies between receive-pack
// implementations in sending status messages.
// TODO: check that server advertises all of these

// this is passed on to git rev-list to enumerate objects, and we're saying don't send the old objects

// TODO: gittuf verify-ref for each dstRef; abort if
// verification fails

// TODO: find better way to evaluate if gittuf refs must
// be pushed

// this is passed on to git rev-list to enumerate objects, and we're saying don't send the old objects

// Write the flush packet as we're done with ref processing

// Write objects that must be pushed to stdin
// the extra \n is used to indicate end of stdin entries

// Redirect packfile bytes to remote service stdin

// Status updates get sent to parent process

// remove length prefix

// replace ng with error

// Trailing newline for end of output

// FIXME: we return in fetch and push when successful, need to assess when
// this is reachable
