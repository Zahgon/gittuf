// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package profile

var stopProfilingQueue = []func() error{}

func StartProfiling(cpuFile, memoryFile string) error { _ = "STUB: not implemented"; return nil }

func StopProfiling() error { _ = "STUB: not implemented"; return nil }
