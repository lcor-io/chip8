package screen

import (
	"math/rand"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/lcor-io/chip8/internal/utils/styles"
)

var (
	fgColor = lipgloss.Color("2")

	emptyPixelStyle = lipgloss.NewStyle()
	fullPixelStyle  = emptyPixelStyle.Foreground(fgColor)
)

type RefreshScreenMsg struct{}

type Tui struct {
	pixels [SCREEN_WIDTH][SCREEN_HEIGHT]Pixel
}

func (s *Tui) Clear() {
	for i := range s.pixels {
		for j := range s.pixels[i] {
			s.pixels[i][j] = false
		}
	}
}

func (s *Tui) GetPixel(x int, y int) Pixel {
	if x >= int(SCREEN_WIDTH) || y >= int(SCREEN_HEIGHT) {
		return false
	}
	return s.pixels[x][y]
}

func (s *Tui) SetPixel(x int, y int, p Pixel) {
	if x >= int(SCREEN_WIDTH) || y >= int(SCREEN_HEIGHT) {
		return
	}
	s.pixels[x][y] = p
}

/**
*
* Bubble tea method implementation
*
**/

func (s *Tui) Init() tea.Cmd {
	return nil
}

func (s *Tui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	for i := range SCREEN_WIDTH {
		for j := range SCREEN_HEIGHT {
			s.SetPixel(int(i), int(j), rand.Intn(2) == 1)
		}
	}

	return s, tea.Batch(cmds...)
}

func (s *Tui) View() string {

	columns := []string{}
	for i := range SCREEN_WIDTH {
		column := []string{}
		for j := range SCREEN_HEIGHT {
			if s.pixels[i][j] {
				column = append(column, fullPixelStyle.Render("██"))
			} else {
				column = append(column, emptyPixelStyle.Render("  "))
			}
		}
		columns = append(columns, lipgloss.JoinVertical(lipgloss.Bottom, column...))
	}

	return styles.Border.Render(lipgloss.JoinHorizontal(lipgloss.Right, columns...))
}
