package screen

import (
	"fmt"
)

type DebugScreen struct {
	pixels [DEFAULT_SCREEN_WIDTH][DEFAULT_SCREEN_HEIGHT]pixel
}

func (s *DebugScreen) Clear() {
	for i := range s.pixels {
		for j := range s.pixels[i] {
			s.pixels[i][j] = false
		}
	}
}

func (s *DebugScreen) SetPixel(x int, y int, p pixel) {
	if x >= int(DEFAULT_SCREEN_WIDTH) || y >= int(DEFAULT_SCREEN_HEIGHT) {
		return
	}
	s.pixels[x][y] = p
}

func (s *DebugScreen) Render() {

	fmt.Println()
	for j := range DEFAULT_SCREEN_HEIGHT {
		for i := range DEFAULT_SCREEN_WIDTH {
			if s.pixels[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}
