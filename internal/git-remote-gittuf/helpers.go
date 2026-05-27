// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"bufio"
	"io"

	"github.com/gittuf/gittuf/experimental/gittuf"
)

type logWriteCloser struct {
	name        string
	writeCloser io.WriteCloser
}

func (l *logWriteCloser) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *logWriteCloser) Close() error { _ = "STUB: not implemented"; return nil }

type logReadCloser struct {
	name       string
	readCloser io.ReadCloser
}

func (l *logReadCloser) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *logReadCloser) Close() error { _ = "STUB: not implemented"; return nil }

type logScanner struct {
	name    string
	scanner *bufio.Scanner
}

func (l *logScanner) Buffer(buf []byte, maxN int) { _ = "STUB: not implemented"; return }

func (l *logScanner) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (l *logScanner) Err() error { _ = "STUB: not implemented"; return nil }

func (l *logScanner) Scan() bool { _ = "STUB: not implemented"; return false }

func (l *logScanner) Split(split bufio.SplitFunc) { _ = "STUB: not implemented"; return }

func (l *logScanner) Text() string { _ = "STUB: not implemented"; return "" }

func log(messages ...any) { _ = "STUB: not implemented"; return }

//nolint:gosec

func packetEncode(str string) []byte { _ = "STUB: not implemented"; return nil }

func getGittufWantsAndHaves(repo *gittuf.Repository, remoteTips map[string]string) (map[string]string, []string, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func getSSHCommand(repo *gittuf.Repository) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func testSSH(sshCmd []string, host string) error { _ = "STUB: not implemented"; return nil }

//nolint:gocritic
//nolint:gosec

// with GitHub, we see exit code 1 while with GitLab and BitBucket,
// we see exit code 0
