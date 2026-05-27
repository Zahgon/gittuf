// Copyright The gittuf Authors
// SPDX-License-Identifier: Apache-2.0

package tui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

const (
	colorRegularText = "#FFFFFF"
	colorFocus       = "#007AFF"
	colorBlur        = "#A0A0A0"
	colorFooter      = "#11ff00"
	colorSubtext     = "#555555"
	colorErrorMsg    = "#FF0000"
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorRegularText)).
			Padding(0, 2).
			MarginTop(1).
			Bold(true)

	itemStyle = lipgloss.NewStyle().
			PaddingLeft(4).
			Foreground(lipgloss.Color(colorRegularText))

	selectedItemStyle = lipgloss.NewStyle().
				PaddingLeft(4).
				Foreground(lipgloss.Color(colorRegularText)).
				Background(lipgloss.Color(colorFocus))

	focusedStyle = lipgloss.NewStyle().
			PaddingLeft(4)

	blurredStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorBlur))

	cursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(colorRegularText))
)

// renderWithMargin wraps content in the standard margin used by all screens.
func renderWithMargin(content string) string { _ = "STUB: not implemented"; return "" }

// renderFooter returns the footer text styled in the footer color.
func renderFooter(text string) string { _ = "STUB: not implemented"; return "" }

// renderErrorMsg returns error messages styled in the error color.
func renderErrorMsg(text string) string { _ = "STUB: not implemented"; return "" }

// renderFormScreen renders a form screen with a title, input fields, help text, and footer.
func (m model) renderFormScreen(title string) string { _ = "STUB: not implemented"; return "" }

// renderListScreen renders a list with help text and footer.
func (m model) renderListScreen(l list.Model, helpText string, emptyMsg string, isEmpty bool) string {
	_ = "STUB: not implemented"
	return ""
}

// screenPolicyRulesHelp returns the help bar for the policy rules view screen.
func screenPolicyRulesHelp(readOnly bool) string { _ = "STUB: not implemented"; return "" }

// screenTrustGlobalRulesHelp returns the help bar for the global rules view screen.
func screenTrustGlobalRulesHelp(readOnly bool) string { _ = "STUB: not implemented"; return "" }

// renderDeleteOverlay renders the delete confirmation prompt.
func renderDeleteOverlay(target string) string { _ = "STUB: not implemented"; return "" }

// View renders the TUI.
func (m model) View() string { _ = "STUB: not implemented"; return "" }
