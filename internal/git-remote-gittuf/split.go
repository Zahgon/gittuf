// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package main

func dropCR(data []byte) []byte { _ = "STUB: not implemented"; return nil }

func splitInput(data []byte, atEOF bool) (int, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Request more data.

// We have the flushPkt that'll otherwise cause it to hang.
// This packet isn't followed by a newline a lot of the time, so we just
// end up requesting data perennially.

// We have more data to process, so we just return the current line.

// Request more data.

func splitOutput(data []byte, atEOF bool) (advance int, token []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// Request more data.

// We have the flushPkt that'll otherwise cause it to hang.
// This packet isn't followed by a newline a lot of the time, so we just
// end up requesting data perennially.

// This is the very last newline, we need to return ErrFinalToken to
// not block the stdin scanner anymore.
// This is the fundamental difference between this function and
// splitInput, because this can block stdin.

// We have more data to process, so we just return the current line.

// Request more data.

func splitPacket(data []byte, atEOF bool) (int, []byte, error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// request more

// request more in a new buffer
