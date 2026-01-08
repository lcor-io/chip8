package emulator

import (
	"github.com/lcor-io/chip8/internal/cpu"
	"github.com/lcor-io/chip8/internal/screen"
)

const DEFAULT_PROCESSOR_FREQUENCY = 250 //Processor frequency in Hertz
const DEFAULT_SCREEN_REFRESH_RATE = 60  //Screen refresh rate in Hertz

type emulator struct {
	Cpu    *cpu.CPU
	Screen screen.Screen
}

func (e *emulator) Render() {
	e.Screen.Render()
}

func Init() *emulator {

	// Initialize screen
	screen := &screen.DebugScreen{}
	screen.Clear()

	// Initialize CPU
	cpu := cpu.New()

	emulator := emulator{
		Cpu:    cpu,
		Screen: screen,
	}
	return &emulator
}
