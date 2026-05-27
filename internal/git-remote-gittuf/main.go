// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"io"
	"os"
)

/*
	Git supports the idea of "remote helpers" which can be used to modify
	interactions with remote repositories.
	https://git-scm.com/docs/gitremote-helpers

	Interactions with a remote take two forms: either we're fetching objects
	(git fetch, pull, clone) or we're sending objects (git push). Git has a
	custom protocol dictating the interactions for fetches and pushes and how
	data is communicated using "packfiles". Note: this protocol is now
	versioned. v0 and v1 are effectively identical except that v1 has an
	explicit version declaration while v0 has nothing.
	https://git-scm.com/docs/pack-protocol

	On the other hand, v2 is a significant departure, and is documented
	separately. Currently, v2 only supports fetches, with pushes using v0/v1.
	https://git-scm.com/docs/protocol-v2

	In both cases, there's an underlying communication protocol. Overwhelmingly,
	this is ssh in the real world, to the point that it is part of the Git
	implementation itself. In contrast, the HTTP(s) / FTP(s) protocols are
	implemented as a remote helper program using curl.

	Setup:
		* git-remote-gittuf must be in PATH.
		* A remote must be configured to use `gittuf::` as the prefix. The
		  result of the remote URL must indicate the underlying transport
		  mechanism. Example: `gittuf::https://github.com/gittuf/gittuf` and
		  `gittuf::git@github.com:gittuf/gittuf`

	Invocation:
	During an interaction with a remote configured using the gittuf:: prefix,
	Git invokes the gittuf helper from the PATH.

	Initial interaction:
	Git learns what capabilities the remote helper implements, and chooses the
	appropriate one for the task at hand. git-remote-gittuf is not consistent in
	the advertised capabilities. This is because when the underlying transport
	is HTTP(s) / FTP(s), we just invoke git-remote-http and relay its
	capabilities back to Git. When the underlying transport is SSH, we advertise
	a lightweight set of capabilities to ensure Git chooses the one we want it
	to for both cases.

	Anatomy of a fetch:
		* Git invokes stateless-connect (protocol v2) to communicate with
		  git-upload-pack on the remote
		* Git invokes ls-refs, a command of git-upload-pack, to learn what refs
		  the remote has and what their tips point to
		  Note: we interpose this to learn the status of gittuf refs on the
		  remote
		* Git negotiates with git-upload-pack the objects it wants based on the
		  refs that must be fetched
		  Note: we interpose this to request gittuf specific objects as well
		* The remote sends a packfile with the requested objects
		* git-remote-gittuf uses update-ref to set the local gittuf refs, as Git
		  will not do this for us

	Anatomy of a push:
		* Git invokes list for-push (protocol v0/v1) to list the refs available
		  on the remote
		* Git invokes push (protocol v0/v1) to indicate what refs must be
		  updated to on the remote
		* A packfile is created and streamed to git-receive-pack on the server
*/

var (
	// logFile is used to debug git-remote-gittuf. It is set using
	// GITTUF_LOG_FILE.
	logFile io.Writer

	// gitVersion contains the version of Git used by the client invoking
	// git-remote-gittuf. It is used to self-identify with a remote service such
	// as git-upload-pack and git-receive-pack.
	gitVersion string

	flushPkt     = []byte{'0', '0', '0', '0'}
	delimiterPkt = []byte{'0', '0', '0', '1'}
	endOfReadPkt = []byte{'0', '0', '0', '2'}
)

const (
	gitUploadPack   = "git-upload-pack"
	gitReceivePack  = "git-receive-pack"
	gittufRefPrefix = "refs/gittuf/"
)

func run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// When cloning/fetching, we have to hang until Git sets things up
// before we can update-ref

// `entries` is sorted by name. The "regular" entries have pack- as a
// prefix. When actually fetching contents, it's stored in a tmp_pack or
// temp_rev file. Therefore, if the last entry starts with pack-, we
// know we don't have a tmp_ file.

// not tmp_pack or tmp_rev

// TODO: this breaks when `git fetch` is invoked explicitly for a gittuf
// ref because Git separately tries to update-ref.
// During wants, check if the latest remote gittuf ref tips are
// requested. Use _only_ the latest so as to avoid any unnecessary blob
// collisions.

//nolint:errcheck

// Uncomment after gittuf can accept a git_dir env var; this will happen
// with the gitinterface PRs naturally.

// TODO: this must either be looped to address each changed ref that
// exists locally or gittuf needs another flag for --all.
// var cmd *exec.Cmd
// if rslTip != "" {
// 	log("we have rsl tip")
// 	cmd = exec.Command("gittuf", "verify-ref", "--from-entry", rslTip, "HEAD")
// } else {
// 	cwd, _ := os.Getwd()
// 	log("we don't have rsl tip", cwd)
// 	cmd = exec.Command("gittuf", "verify-ref", "HEAD")
// }
// _, err := cmd.Output()
// if err != nil {
// 	log(err.Error())
// 	if _, nerr := os.Stderr.Write([]byte("gittuf verification failed\n")); nerr != nil {
// 		return errors.Join(err, nerr)
// 	}
// 	return err
// }

func populateGitVersion() error { _ = "STUB: not implemented"; return nil }

func main() {
	logFilePath := os.Getenv("GITTUF_LOG_FILE")
	if logFilePath != "" {
		file, err := os.Create(logFilePath) //nolint:gosec
		if err != nil {
			panic(err)
		}

		logFile = file
	}

	if err := populateGitVersion(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	if err := run(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
