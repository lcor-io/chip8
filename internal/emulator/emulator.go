package emulator

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lcor-io/chip8/internal/cpu"
	"github.com/lcor-io/chip8/internal/screen"
	sc "github.com/lcor-io/chip8/internal/screen"
)

const DEFAULT_PROCESSOR_FREQUENCY = 250 //Processor frequency in Hertz
const DEFAULT_SCREEN_REFRESH_RATE = 60  //Screen refresh rate in Hertz

type emulator struct {
	Cpu    *cpu.CPU
	Screen screen.Screen

	procTicker   time.Ticker
	screenTicker time.Ticker
}

func (e *emulator) TUILoop(p *tea.Program) {
	for {
		select {
		case <-e.screenTicker.C:

			// Random pixels for testing
			for i := range 64 {
				for j := range 32 {
					e.Screen.SetPixel(i, j, rand.Intn(2) == 1)
				}
			}
			p.Send(sc.RefreshScreenMsg{})
		case <-e.procTicker.C:
			//TODO get actions from proc
		}
	}
}

func Init() *emulator {

	// Initialize screen
	screen := &sc.TUIScreen{}
	screen.Clear()

	// Initialize CPU
	cpu := cpu.New()

	emulator := emulator{
		Cpu:    cpu,
		Screen: screen,

		procTicker:   *time.NewTicker(time.Second / DEFAULT_PROCESSOR_FREQUENCY),
		screenTicker: *time.NewTicker(time.Second / DEFAULT_SCREEN_REFRESH_RATE),
	}

	// Start in TUI Mode
	p := tea.NewProgram(screen, tea.WithAltScreen())
	go emulator.TUILoop(p)
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occured at launch")
		os.Exit(1)
	}

	return &emulator
}
