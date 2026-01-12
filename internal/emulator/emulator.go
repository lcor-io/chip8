package emulator

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lcor-io/chip8/internal/cpu"
	"github.com/lcor-io/chip8/internal/keyboard"
	"github.com/lcor-io/chip8/internal/screen"
)

const DEFAULT_SCREEN_REFRESH_RATE = 60                              //In Hertz
const DEFAULT_PROCESSOR_FREQUENCY = DEFAULT_SCREEN_REFRESH_RATE * 4 //In Hertz

type emulator struct {
	Cpu      *cpu.CPU
	Screen   screen.Screen
	Keyboard *keyboard.Keyboard

	procTicker   *time.Ticker
	screenTicker *time.Ticker

	running bool

	t_width  int
	t_height int
	t_ready  bool
}

func (e *emulator) TuiEventLoop(p *tea.Program) {
	for range e.screenTicker.C {
		p.Send(screen.RefreshScreenMsg{})
	}
}

func (e *emulator) cpuLoop() {
	for range e.procTicker.C {
		if e.running {
			e.interpret()
		}
	}
}

func (e *emulator) interpret() {
	instruction := e.Cpu.GetInstruction() // Get instruction from the CPU

	// Get important bytes
	b1 := uint8((instruction & (0x0F00)) >> 8)
	b2 := uint8((instruction & (0x00F0)) >> 4)
	b3 := uint8(instruction & 0x000F)

	opcode, _ := cpu.GetOpcode(instruction) // Get the opcode from the instruction

	switch opcode {
	case cpu.OPCODE_0NNN:
		//TODO Special instruction, not used in this emulator
	case cpu.OPCODE_00E0:
		e.opcode_00E0()
	case cpu.OPCODE_00EE:
		e.opcode_00EE()
	case cpu.OPCODE_1NNN:
		e.opcode_1NNN(b1, b2, b3)
	case cpu.OPCODE_2NNN:
		e.opcode_2NNN(b1, b2, b3)
	case cpu.OPCODE_3XNN:
		e.opcode_3XNN(b1, b2, b3)
	case cpu.OPCODE_4XNN:
		e.opcode_4XNN(b1, b2, b3)
	case cpu.OPCODE_5XY0:
		e.opcode_5XY0(b1, b2)
	case cpu.OPCODE_6XNN:
		e.opcode_6XNN(b1, b2, b3)
	case cpu.OPCODE_7XNN:
		e.opcode_7XNN(b1, b2, b3)
	case cpu.OPCODE_8XY0:
		e.opcode_8XY0(b1, b2)
	case cpu.OPCODE_8XY1:
		e.opcode_8XY1(b1, b2)
	case cpu.OPCODE_8XY2:
		e.opcode_8XY2(b1, b2)
	case cpu.OPCODE_8XY3:
		e.opcode_8XY3(b1, b2)
	case cpu.OPCODE_8XY4:
		e.opcode_8XY4(b1, b2)
	case cpu.OPCODE_8XY5:
		e.opcode_8XY5(b1, b2)
	case cpu.OPCODE_8XY6:
		e.opcode_8XY6(b1)
	case cpu.OPCODE_8XY7:
		e.opcode_8XY7(b1, b2)
	case cpu.OPCODE_8XYE:
		e.opcode_8XYE(b1)
	case cpu.OPCODE_9XY0:
		e.opcode_9XY0(b1, b2)
	case cpu.OPCODE_ANNN:
		e.opcode_ANNN(b1, b2, b3)
	case cpu.OPCODE_BNNN:
		e.opcode_BNNN(b1, b2, b3)
	case cpu.OPCODE_CXNN:
		e.opcode_CXNN(b1, b2, b3)
	case cpu.OPCODE_DXYN:
		e.opcode_DXYN(b1, b2, b3)
	case cpu.OPCODE_EX9E:
		e.opcode_EX9E(b1)
	case cpu.OPCODE_EXA1:
		e.opcode_EXA1(b1)
	case cpu.OPCODE_FX07:
		e.opcode_FX07(b1)
	case cpu.OPCODE_FX0A:
		e.opcode_FX0A(b1)
	case cpu.OPCODE_FX15:
		e.opcode_FX15(b1)
	case cpu.OPCODE_FX18:
		e.opcode_FX18(b1)
	case cpu.OPCODE_FX1E:
		e.opcode_FX1E(b1)
	case cpu.OPCODE_FX29:
		e.opcode_FX29(b1)
	case cpu.OPCODE_FX33:
		e.opcode_FX33(b1)
	case cpu.OPCODE_FX55:
		e.opcode_FX55(b1)
	case cpu.OPCODE_FX65:
		e.opcode_FX65(b1)
	}
}

func (e *emulator) Start() {
	e.running = true
}
func (e *emulator) Stop() {
	e.running = false
}
func (e *emulator) Restart() {
	e.running = false
	e.Screen.Clear()
	e.Cpu.Pc = cpu.INITIAL_RAM_ADDRESS
	e.running = true
}
func (e *emulator) LoadROM(path string) error {
	return e.Cpu.LoadRom(path)
}

func Init() *emulator {

	// Initialize Keyboard
	keyboard := keyboard.New()

	// Initialize CPU
	cpu := cpu.New()

	// Initialize screen
	screen := &screen.Tui{}
	screen.Clear()

	e := emulator{
		Cpu:      cpu,
		Screen:   screen,
		Keyboard: keyboard,

		procTicker:   time.NewTicker(time.Second / DEFAULT_PROCESSOR_FREQUENCY),
		screenTicker: time.NewTicker(time.Second / DEFAULT_SCREEN_REFRESH_RATE),

		running: false,
	}

	// Launch the CPU
	go e.cpuLoop()

	return &e
}
