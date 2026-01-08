package screen

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	fgColor = lipgloss.Color("1")

	emptyPixelStyle = lipgloss.NewStyle()
	fullPixelStyle  = emptyPixelStyle.Foreground(fgColor)
)

type RefreshScreenMsg struct{}

type TUIScreen struct {
	pixels [DEFAULT_SCREEN_WIDTH][DEFAULT_SCREEN_HEIGHT]pixel

	t_width  int
	t_height int
}

func (s *TUIScreen) Clear() {
	for i := range s.pixels {
		for j := range s.pixels[i] {
			s.pixels[i][j] = false
		}
	}
}

func (s *TUIScreen) SetPixel(x int, y int, p pixel) {
	if x >= int(DEFAULT_SCREEN_WIDTH) || y >= int(DEFAULT_SCREEN_HEIGHT) {
		return
	}
	s.pixels[x][y] = p
}

func (s *TUIScreen) Render() string {
	columns := []string{}
	for i := range DEFAULT_SCREEN_WIDTH {
		column := []string{}
		for j := range DEFAULT_SCREEN_HEIGHT {
			if s.pixels[i][j] {
				column = append(column, fullPixelStyle.Render("██"))
			} else {
				column = append(column, emptyPixelStyle.Render("  "))
			}
		}
		columns = append(columns, lipgloss.JoinVertical(lipgloss.Bottom, column...))
	}

	screen := lipgloss.JoinHorizontal(lipgloss.Right, columns...)

	if s.t_width > int(DEFAULT_SCREEN_WIDTH) && s.t_height > int(DEFAULT_SCREEN_HEIGHT) {
		return lipgloss.Place(s.t_width, s.t_height, lipgloss.Center, lipgloss.Center, screen)
	}

	return screen
}

/**
*
* Bubble tea method implementation
*
**/

func (s *TUIScreen) Init() tea.Cmd {
	return nil
}

func (s *TUIScreen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return s, tea.Quit
		}
	case tea.WindowSizeMsg:
		s.t_width = msg.Width
		s.t_height = msg.Height
	case RefreshScreenMsg:
	}

	return s, nil
}

func (s *TUIScreen) View() string {
	return s.Render()
}
