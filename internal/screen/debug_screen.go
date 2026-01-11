package screen

import (
	"fmt"
)

type DebugScreen struct {
	pixels [SCREEN_WIDTH][SCREEN_HEIGHT]Pixel
}

func (s *DebugScreen) Clear() {
	for i := range s.pixels {
		for j := range s.pixels[i] {
			s.pixels[i][j] = false
		}
	}
}

func (s *DebugScreen) SetPixel(x int, y int, p Pixel) {
	if x >= int(SCREEN_WIDTH) || y >= int(SCREEN_HEIGHT) {
		return
	}
	s.pixels[x][y] = p
}

func (s *DebugScreen) Render() {

	fmt.Println()
	for j := range SCREEN_HEIGHT {
		for i := range SCREEN_WIDTH {
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
