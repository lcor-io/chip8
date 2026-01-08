package main

import "github.com/lcor-io/chip8/internal/emulator"

func main() {

	emulator := emulator.Init()

	for i := range 64 {
		for j := range 32 {
			emulator.Screen.SetPixel(i, j, ((i % (j + 1)) != 0))
		}
	}

	emulator.Screen.Render()

}
