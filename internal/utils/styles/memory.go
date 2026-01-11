package styles

import "github.com/charmbracelet/lipgloss"

var (
	PCMemoryCell = lipgloss.NewStyle().
			Background(FgColor)

	RegisterTitle = lipgloss.NewStyle().
			Foreground(FgColor).
			MarginTop(1).
			Bold(true)
	RegisterListEnumerator = lipgloss.NewStyle().
				Margin(0, 1).
				Bold(true)
	RegisterListItem = lipgloss.NewStyle()
)
