package styles

import "github.com/charmbracelet/lipgloss"

var (
	KeyboardCell = lipgloss.NewStyle().
			Align(lipgloss.Center, lipgloss.Center).
			Margin(0).
			Border(lipgloss.NormalBorder()).
			BorderForeground(FgColor)

	KeyboardPressedCell = KeyboardCell.
				Align(lipgloss.Center, lipgloss.Center).
				Background(FgColor).
				Border(lipgloss.NormalBorder()).
				BorderForeground(FgColor)
)
