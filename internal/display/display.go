// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package display

import (
	"io"
	"os/exec"
)

func NewDisplayWriter(output io.Writer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// adityasaky: os.Stdout is an io.WriteCloser and we hardcode that as our
// output medium. So, do we even need this check and noopwritecloser?
// Possibly not, but I suggest we keep it until we can sufficiently evaluate
// across multiple environments. noopwritecloser is handy for test writers
// as well, so.

// pagerGetter is a pointer to the dispatcher that selects the pager program to
// use. We use this to override the "real" pagerGetter in tests.
type pagerGetter = func() pager

var getPager pagerGetter = getPagerReal //nolint:revive

func getPagerReal() pager { _ = "STUB: not implemented"; return *new(pager) }

// look at what the user has configured
// default on unix-like systems
// default on Windows

// pager implements a generic interface for a stdout pager like less or more. We
// use this interface because some environments need coloring information to be
// explicitly enabled with the pager.
type pager interface {
	getBinary() string
	getFlags() []string
}

// pagerEnvVar inspects the user's $PAGER variable and creates a pager instance
// using the information there.
type pagerEnvVar struct {
	binary string
	flags  []string
}

func newPagerEnvVar() *pagerEnvVar { _ = "STUB: not implemented"; return nil }

// look at what the user has configured

func (p *pagerEnvVar) getBinary() string { _ = "STUB: not implemented"; return "" }

func (p *pagerEnvVar) getFlags() []string {
	_ = "STUB: not implemented"

	// pagerLess implements the pager interface for `less`.
	return nil
}

type pagerLess struct{}

func newPagerLess() *pagerLess { _ = "STUB: not implemented"; return nil }

func (p *pagerLess) getBinary() string { _ = "STUB: not implemented"; return "" }

func (p *pagerLess) getFlags() []string { _ = "STUB: not implemented"; return nil }

// see if user already has preferred LESS flags

// we prefix the "-" as this is passed to exec.Command

// These are the default flags to use with less, matches what Git sets

// --quit-if-one-screen
// --RAW-CONTROL-CHARS for coloring
// --no-init

// pagerMore implements the pager interface for `more`.
type pagerMore struct{}

func newPagerMore() *pagerMore { _ = "STUB: not implemented"; return nil }

func (p *pagerMore) getBinary() string { _ = "STUB: not implemented"; return "" }

func (p *pagerMore) getFlags() []string {
	_ = "STUB: not implemented"

	// pagerWriteCloser implements the io.WriteCloser while supporting writing
	// buffered contents displayed using a pager program like less or more.
	return nil
}

type pagerWriteCloser struct {
	command     *exec.Cmd
	stdInWriter io.WriteCloser
	started     bool
}

func (p *pagerWriteCloser) Write(contents []byte) (int, error) {
	_ = "STUB: not implemented"

	// Load the page program's stdin pipe so we can feed it content to
	// display
	return 0, nil
}

// Start the page cmd

func (p *pagerWriteCloser) Close() error {
	_ = "STUB: not implemented"
	// Close the stdin pipe first as the cmd will wait indefinitely otherwise
	return nil
}

// noopwritecloser is a fallback to convert an io.Writer into io.WriteCloser. It
// adds a Close method which does nothing (i.e., it's a noop).
type noopwritecloser struct {
	writer io.Writer
}

func (n *noopwritecloser) Write(contents []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (n *noopwritecloser) Close() error { _ = "STUB: not implemented"; return nil }
