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

func (e *emulator) tuiLoop(p *tea.Program) {
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
			e.interpret()
		}
	}
}

func (e *emulator) interpret() {
	instruction := e.Cpu.GetInstruction()
	b3 := uint8((instruction & (0x0F00)) >> 8)
	b2 := uint8((instruction & (0x00F0)) >> 4)
	b1 := uint8(instruction & 0x000F)

	opcode, _ := cpu.GetOpcode(instruction)

	switch opcode {
	case cpu.OPCODE_0NNN:
	case cpu.OPCODE_00E0:
	case cpu.OPCODE_00EE:
	case cpu.OPCODE_1NNN:
	case cpu.OPCODE_2NNN:
	case cpu.OPCODE_3XNN:
	case cpu.OPCODE_4XNN:
	case cpu.OPCODE_5XY0:
	case cpu.OPCODE_6XNN:
	case cpu.OPCODE_7XNN:
	case cpu.OPCODE_8XY0:
	case cpu.OPCODE_8XY1:
	case cpu.OPCODE_8XY2:
	case cpu.OPCODE_8XY3:
	case cpu.OPCODE_8XY4:
	case cpu.OPCODE_8XY5:
	case cpu.OPCODE_8XY6:
	case cpu.OPCODE_8XY7:
	case cpu.OPCODE_8XYE:
	case cpu.OPCODE_9XY0:
	case cpu.OPCODE_ANNN:
	case cpu.OPCODE_BNNN:
	case cpu.OPCODE_CXNN:
	case cpu.OPCODE_DXYN:
		e.opcode_DXYN(b1, b2, b3)
	case cpu.OPCODE_EX9E:
	case cpu.OPCODE_EXA1:
	case cpu.OPCODE_FX07:
	case cpu.OPCODE_FX0A:
	case cpu.OPCODE_FX15:
	case cpu.OPCODE_FX18:
	case cpu.OPCODE_FX1E:
	case cpu.OPCODE_FX29:
	case cpu.OPCODE_FX33:
	case cpu.OPCODE_FX55:
	case cpu.OPCODE_FX65:
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
	go emulator.tuiLoop(p)
	if _, err := p.Run(); err != nil {
		fmt.Printf("An error occured at launch")
		os.Exit(1)
	}

	return &emulator
}
