// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package version //nolint:revive

// gitVersion records the basic version information from Git. It is typically
// overwritten during a go build.
var gitVersion = "devel"

func GetVersion() string { _ = "STUB: not implemented"; return "" }
