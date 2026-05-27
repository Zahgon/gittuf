// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update updates the model based on the message received.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// Stay on the loading screen so the error is visible; user can press q to quit.

// Delete confirmation overlay intercepts all keys

// Global handlers (quit, back navigation)

// Only quit from non-form screens (avoid consuming 'q' in text inputs)

// Screen-specific input handling

// Delegate to active bubbles component per screen

// handleEnter handles the enter key press on selection menu screens.
func (m model) handleEnter() (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// handleRulesListKey handles keybindings on rule list screens (rules and global rules).
// For unhandled keys (including up/down arrows), it delegates to the active list for navigation.
func (m model) handleRulesListKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// add rule

// edit rule

// delete rule

// reorder up

// reorder down

// Delegate unhandled keys to the active list for navigation (up/down arrows, etc.)

// handleDeleteConfirm handles the delete confirmation overlay.
func (m model) handleDeleteConfirm(key string) (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// handlePolicyFormSubmit handles enter on policy add/edit form screens.
func (m model) handlePolicyFormSubmit() (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// Not on last field yet - just advance focus

// handleGlobalFormSubmit handles enter on global add/edit form screens.
func (m model) handleGlobalFormSubmit() (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// handleReorderUp moves the selected rule up in the list.
func (m model) handleReorderUp() (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// handleReorderDown moves the selected rule down in the list.
func (m model) handleReorderDown() (tea.Model, tea.Cmd) {
	_ = "STUB: not implemented"
	return *new(tea.Model), *new(tea.Cmd)
}

// cycleFocus moves focus (the cursor) between input fields in form screens.
func (m *model) cycleFocus(key string) { _ = "STUB: not implemented"; return }

// splitAndTrim splits a comma-separated string and trims whitespace.
func splitAndTrim(s string) []string { _ = "STUB: not implemented"; return nil }
