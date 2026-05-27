// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"errors"

	"github.com/gittuf/gittuf/experimental/gittuf"
)

var ErrFailedAuthentication = errors.New("failed getting remote refs")

// handleCurl implements the helper for remotes configured to use the curl
// backend. For this transport, we invoke git-remote-http, only interjecting at
// specific points to make gittuf specific additions.
func handleCurl(ctx context.Context, repo *gittuf.Repository, remoteName, url string) (map[string]string, bool, error) {
	_ = "STUB: not implemented"
	// Scan git-remote-gittuf stdin for commands from the parent process
	return nil, false, nil
}

// We invoke git-remote-http, itself a Git remote helper
//nolint:gosec

// We want to inspect the helper's stdout for the gittuf ref statuses

// We want to interpose with the helper's stdin to push and fetch gittuf
// specific objects and refs

/*
	stateless-connect is the new experimental way of
	communicating with the remote. It implements Git Protocol
	v2. Here, we don't do much other than recognizing that we're
	in a fetch, as this protocol doesn't support pushes yet.
*/

// Write to git-remote-http

// Receive the initial info sent by the service via
// git-remote-http

// We wrap this in an extra loop because for
// some reason, on Windows, git-remote-http
// responds with a buffer that just contains
// `\n` followed by the actual response.
// However, the initial buffer is taken to be
// the end of output, meaning we miss the actual
// end of output.

// If nothing is returned, the user has likely failed to
// authenticate with the remote

// flushPkt is used to indicate the end of
// output

// Read in command from parent process -> this should be
// command=ls-refs with protocol v2
// ls-refs is a command to upload-pack. Like list and list
// for-push, it enumerates the refs and their states on the remote.
// Unlike those commands, this must be passed to upload-pack.
// Further, ls-refs can be parametrized with ref-prefixes. We add
// refs/gittuf/ as a prefix to learn about the gittuf refs on the
// remote during fetches, but only if the client already has
// ref-prefixes set.

// Add ref-prefix refs/gittuf/ to the ls-refs command before
// flush, but only if the client sent ref-prefixes

// flushPkt is used to indicate the end of input

// Read advertised refs from the remote

// remove pkt length prefix

// If the gittuf ref is the very first, then there will be
// additional information in the output after a null byte.
// However, this is unlikely as HEAD is typically the first.

// drop everything from null byte onwards

// Write output to parent process

// endOfReadPkt indicates end of response
// in stateless connections

// At this point, we enter the haves / wants negotiation, which is
// followed usually by the remote sending a packfile with the
// requested Git objects.

// We add the gittuf specific objects as wants. We don't have to
// specify haves as Git automatically specifies all the objects it
// has regardless of what refs they're reachable via.

// Read in command from parent process -> this should be
// command=fetch with protocol v2

// track this in case there are multiple rounds of negotiation

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

// On a clone, we see `done` and
// then flush. We need to write
// our wants before done, but
// wroteWants can't be set to
// true until the next buffer
// with flush is written to the
// remote.

// Take out this ref as
// something for us to
// update or add wants
// for

// Read from remote if wants are done
// We may need to scan multiple times for inputs, which is
// why this flag is used

// TODO: check multiplexed output

// Send along to parent process

// Two things are possible
// a) The communication is done
// b) The remote indicates another round of
// negotiation is required
// Instead of parsing the output to find out,
// we let the parent process tell us
// If the parent process has further input, more
// negotiation is needed

// Take out this ref as
// something for us to
// update or add wants
// for

// Having scanned already, we must write prior
// to letting the scan continue in the outer
// loop
// This assumes the very first input isn't just
// flush again...

/*
	The helper has two commands, in reality: list, list
	for-push. Both of these are used to list the states of refs
	on the remote. The for-push variation just formats it in a
	way that can be used for the push command later.

	We inspect this to learn we're in a push. We also use the
	output of this command, implemented by git-remote-https, to
	learn what the states of the gittuf refs are on the remote.
*/

// Write it to git-remote-http

// Read remote refs

// If nothing is returned, the user has likely failed to
// authenticate with the remote

// Inspect each one to see if it's a gittuf ref

// Pass remote ref status to parent process

// flushPkt indicates end of message

// multiline input

// dstRefs tracks the explicitly pushed refs so we know
// to pass the response from the server for those refs
// back to Git

// TODO: maybe find another way to determine
// whether repo is gittuf enabled
// The remote may not have gittuf refs but the
// local may, meaning this won't get synced

// force push
// TODO: during a force push, we want to also revoke prior
// pushes

// Create RSL entries for the ref as long as it's not a
// gittuf ref
// A gittuf ref can pop up here when it's explicitly
// pushed by the user

// TODO: skipping propagation; invoke it once total instead of per ref

// Write push command to helper

// Push RSL if it hasn't been explicitly pushed

// Indicate end of push statements

// We wrap this in an extra loop because for
// some reason, on Windows, the trailing newline
// indicating end of output is sent in a
// separate buffer that's otherwise missed. If
// we miss that newline, we hang as though the
// push isn't complete.

// outputSplit has either two items or
// three items. It has two when the
// response is `ok` and potentially
// three when the response is `error`.
// Either way, the second item is the
// ref in question that we want to
// bubble back to our caller.

// This should never happen but
// if it does, just send it back
// to the caller

// this was explicitly
// pushed by the user

// Pass through other commands we don't want to interpose to the
// curl helper

// Receive the initial info sent by the service

// Check for end of message
