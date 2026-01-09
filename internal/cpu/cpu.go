package cpu

const (
	RAM_SIZE            uint16 = 4_096 //Ram size is set to 4kb
	INITIAL_RAM_ADDRESS uint16 = 0x200 //Inital address at CPU startup

	REGISTERS_NUMBER uint8 = 16 //There are 16 registers on the CPU
	STACK_SIZE       uint8 = 16 //Stack has 16 levels
)

type CPU struct {
	Memory    [RAM_SIZE]uint8 //Internal memory
	Memory_pc uint16          //RAM counter

	V [REGISTERS_NUMBER]uint8 //Registers
	I uint16                  //Memory register

	Stack    [STACK_SIZE]uint16
	Stack_pc int16

	Sys_counter   uint8 //System counter used for game events. Can be set and read
	Sound_counter uint8 //Counter for sound effect. When value is not zero, a sound signal is emited
}

func New() *CPU {

	cpu := CPU{}

	// Initialize RAM
	for i := range cpu.Memory {
		cpu.Memory[i] = 0x0
	}
	cpu.Memory_pc = INITIAL_RAM_ADDRESS

	return &cpu
}
