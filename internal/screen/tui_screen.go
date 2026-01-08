package screen

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	bgColor = lipgloss.Color("99")

	emptyPixelStyle = lipgloss.NewStyle().Height(1).Width(1)
	fullPixelStyle  = emptyPixelStyle.Background(bgColor)
)

type TUIScreen struct {
	pixels [DEFAULT_SCREEN_WIDTH][DEFAULT_SCREEN_HEIGHT]pixel
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

func (s *TUIScreen) Render() {

	columns := []string{}
	for i := range DEFAULT_SCREEN_WIDTH {
		column := []string{}
		for j := range DEFAULT_SCREEN_HEIGHT {
			if s.pixels[i][j] {
				column = append(column, fullPixelStyle.Render())
			} else {
				column = append(column, emptyPixelStyle.Render())
			}
		}
		columns = append(columns, lipgloss.JoinVertical(lipgloss.Bottom, column...))
	}

	screen := lipgloss.JoinHorizontal(lipgloss.Right, columns...)
	fmt.Println(screen)
}
