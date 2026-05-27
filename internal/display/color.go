// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package display

type color uint

func (c color) Code() string { _ = "STUB: not implemented"; return "" }

const (
	reset color = iota
	red
	green
	yellow
	blue
	magenta
	cyan
	gray
	white
)

type colorerFunc = func(string, color) string

var colorer colorerFunc = colorerOn //nolint:revive

func colorerOn(s string, c color) string { _ = "STUB: not implemented"; return "" }

func colorerOff(s string, _ color) string { _ = "STUB: not implemented"; return "" }

func EnableColor() { _ = "STUB: not implemented"; return }

func DisableColor() { _ = "STUB: not implemented"; return }
