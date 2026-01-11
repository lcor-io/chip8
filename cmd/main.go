package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	emu "github.com/lcor-io/chip8/internal/emulator"
)

func main() {

	emulator := emu.Init()
	emulator.LoadROM("roms/space_invaders.ch8")

	// Start in TUI Mode
	p := tea.NewProgram(emulator, tea.WithAltScreen(), tea.WithMouseCellMotion())
	go emulator.TuiEventLoop(p)
	emulator.Start()

	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occured at launch")
		os.Exit(1)
	}
}
