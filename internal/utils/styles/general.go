package styles

import "github.com/charmbracelet/lipgloss"

var (
	FgColor = lipgloss.Color("2")

	Border = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(FgColor)
)
