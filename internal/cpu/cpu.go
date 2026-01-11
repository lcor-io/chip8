package cpu

import (
	"time"

	"github.com/charmbracelet/bubbles/viewport"
)

const (
	RAM_SIZE            uint16 = 4_096 //Ram size is set to 4kb
	INITIAL_RAM_ADDRESS uint16 = 0x200 //Inital address at CPU startup

	REGISTERS_NUMBER uint8 = 16 //There are 16 registers on the CPU
	STACK_SIZE       uint8 = 16 //Stack has 16 levels

	DEFAULT_INTERNAL_COUNTER_FREQUENCY = 60 //In Hertz
)

type CPU struct {
	Memory [RAM_SIZE]uint8 //Internal memory
	Pc     uint16          //RAM counter

	V [REGISTERS_NUMBER]uint8 //Registers
	I uint16                  //Memory register

	Stack [STACK_SIZE]uint16
	Sc    uint8

	Delay_timer   uint8 //System counter used for game events. Can be set and read
	Sound_timer   uint8 //Counter for sound effect. When value is not zero, a sound signal is emited
	signalsTicker time.Ticker

	// below is used for tui rendering
	ready    bool
	viewport viewport.Model
}

// Signal loop is decreasing internal counters
func (c *CPU) signalLoop() {

	for range c.signalsTicker.C {
		if c.Sound_timer > 0 {
			c.Sound_timer--
		}
		if c.Delay_timer > 0 {
			c.Delay_timer--
		}

	}

}

// This method is called to store font in hexadecimal formats in unused memory (from 0x000 to 0x200)
func (c *CPU) initializeFonts() {
	// "0"
	c.Memory[0] = 0xF0
	c.Memory[1] = 0x90
	c.Memory[2] = 0x90
	c.Memory[3] = 0x90
	c.Memory[4] = 0xF0
	// "1"
	c.Memory[5] = 0x20
	c.Memory[6] = 0x60
	c.Memory[7] = 0x20
	c.Memory[8] = 0x20
	c.Memory[9] = 0x70
	// "2"
	c.Memory[10] = 0xF0
	c.Memory[11] = 0x10
	c.Memory[12] = 0xF0
	c.Memory[13] = 0x80
	c.Memory[14] = 0xF0
	// "3"
	c.Memory[15] = 0xF0
	c.Memory[16] = 0x10
	c.Memory[17] = 0xF0
	c.Memory[18] = 0x10
	c.Memory[19] = 0xF0
	// "4"
	c.Memory[20] = 0x90
	c.Memory[21] = 0x90
	c.Memory[22] = 0xF0
	c.Memory[23] = 0x10
	c.Memory[24] = 0x10
	// "5"
	c.Memory[25] = 0xF0
	c.Memory[26] = 0x80
	c.Memory[27] = 0xF0
	c.Memory[28] = 0x10
	c.Memory[29] = 0xF0
	// "6"
	c.Memory[30] = 0xF0
	c.Memory[31] = 0x80
	c.Memory[32] = 0xF0
	c.Memory[33] = 0x90
	c.Memory[34] = 0xF0
	// "7"
	c.Memory[35] = 0xF0
	c.Memory[36] = 0x10
	c.Memory[37] = 0x20
	c.Memory[38] = 0x40
	c.Memory[39] = 0x40
	// "8"
	c.Memory[40] = 0xF0
	c.Memory[41] = 0x90
	c.Memory[42] = 0xF0
	c.Memory[43] = 0x90
	c.Memory[44] = 0xF0
	// "9"
	c.Memory[45] = 0xF0
	c.Memory[46] = 0x90
	c.Memory[47] = 0xF0
	c.Memory[48] = 0x10
	c.Memory[49] = 0xF0
	// "A"
	c.Memory[50] = 0xF0
	c.Memory[51] = 0x90
	c.Memory[52] = 0xF0
	c.Memory[53] = 0x90
	c.Memory[54] = 0x90
	// "B"
	c.Memory[55] = 0xE0
	c.Memory[56] = 0x90
	c.Memory[57] = 0xE0
	c.Memory[58] = 0x10
	c.Memory[59] = 0xE0
	// "C"
	c.Memory[60] = 0xF0
	c.Memory[61] = 0x80
	c.Memory[62] = 0x80
	c.Memory[63] = 0x80
	c.Memory[64] = 0xF0
	// "D"
	c.Memory[65] = 0xE0
	c.Memory[66] = 0x90
	c.Memory[67] = 0x90
	c.Memory[68] = 0x90
	c.Memory[69] = 0xE0
	// "E"
	c.Memory[70] = 0xF0
	c.Memory[71] = 0x80
	c.Memory[72] = 0xF0
	c.Memory[73] = 0x80
	c.Memory[74] = 0xF0
	// "F"
	c.Memory[75] = 0xF0
	c.Memory[76] = 0x80
	c.Memory[77] = 0xF0
	c.Memory[78] = 0x80
	c.Memory[79] = 0x80
}

func New() *CPU {

	cpu := CPU{}

	// Initialize RAM
	for i := range cpu.Memory {
		cpu.Memory[i] = 0x0
	}
	cpu.Pc = INITIAL_RAM_ADDRESS

	// Initialize fonts in memory
	cpu.initializeFonts()

	// Launch the internal signal loop
	cpu.Delay_timer = 0
	cpu.Sound_timer = 0
	go cpu.signalLoop()

	return &cpu
}
