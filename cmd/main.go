package main

import (
	emu "github.com/lcor-io/chip8/internal/emulator"
)

func main() {

	emulator := emu.Init()
	emulator.LoadROM("../roms/maze.ch8")
	emulator.Start()

}
