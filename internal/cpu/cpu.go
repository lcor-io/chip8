package cpu

import "time"

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

func New() *CPU {

	cpu := CPU{}

	// Initialize RAM
	for i := range cpu.Memory {
		cpu.Memory[i] = 0x0
	}
	cpu.Pc = INITIAL_RAM_ADDRESS

	// Launch the internal signal loop
	cpu.Delay_timer = 0
	cpu.Sound_timer = 0
	go cpu.signalLoop()

	return &cpu
}
