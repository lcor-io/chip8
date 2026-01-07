package cpu

const (
	RAM_SIZE            uint16 = 4_096 //Ram size is set to 4kb
	INITIAL_RAM_ADDRESS uint16 = 0x200 //Inital address at cpu startup

	REGISTERS_NUMBER uint8 = 16 //There are 16 registers on the CPU
	STACK_SIZE       uint8 = 16 //Stack has 16 levels
)

type cpu struct {
	ram    [RAM_SIZE]uint8 //Internal memory
	ram_pc uint16          //RAM counter

	reg         [REGISTERS_NUMBER]uint8 //Registers
	reg_address uint16                  //Memory register

	stack    [STACK_SIZE]uint16
	stack_pc int16

	sys_counter   uint8 //System counter used for game events. Can be set and read
	sound_counter uint8 //Counter for sound effect. When value is not zero, a sound signal is emited
}

// This method is used for decrementing the cpu counters. The system run at 60Hz, and thus this method should be run every 16ms
func (c *cpu) Count() {
	if c.sys_counter > 0 {
		c.sys_counter--
	}
	if c.sound_counter > 0 {
		c.sound_counter--
	}
}

func New() *cpu {

	cpu := cpu{}

	// Initialize RAM
	for i := range cpu.ram {
		cpu.ram[i] = 0x0
	}
	cpu.ram_pc = INITIAL_RAM_ADDRESS

	return &cpu
}
