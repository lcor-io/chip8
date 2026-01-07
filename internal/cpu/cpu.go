package cpu

import "os"

const (
	RAM_SIZE            uint16 = 4_096 //Ram size is set to 4kb
	INITIAL_RAM_ADDRESS uint16 = 0x200 //Inital address at cpu startup
	REGISTERS_NUMBER    uint8  = 16    //There are 16 registers on the CPU
	STACK_SIZE          uint8  = 16
)

type cpu struct {
	ram    [RAM_SIZE]uint8 //Internal memory
	ram_pc uint16          //Program counter

	reg         [REGISTERS_NUMBER]uint8 //Registers
	reg_address uint16                  //Memory register

	stack    [STACK_SIZE]uint16
	stack_pc int16

	sys_counter   uint8
	sound_counter uint8
}


func New() *cpu {

	cpu := cpu{}
	cpu.ram_pc = INITIAL_RAM_ADDRESS

	return &cpu
}
