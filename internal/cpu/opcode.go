package cpu

import "fmt"

type opcode uint8

// The list of all 35 instructions from the chip8 processor
const (
	OPCODE_0NNN opcode = iota
	OPCODE_00E0
	OPCODE_00EE
	OPCODE_1NNN
	OPCODE_2NNN
	OPCODE_3XNN
	OPCODE_4XNN
	OPCODE_5XY0
	OPCODE_6XNN
	OPCODE_7XNN
	OPCODE_8XY0
	OPCODE_8XY1
	OPCODE_8XY2
	OPCODE_8XY3
	OPCODE_8XY4
	OPCODE_8XY5
	OPCODE_8XY6
	OPCODE_8XY7
	OPCODE_8XYE
	OPCODE_9XY0
	OPCODE_ANNN
	OPCODE_BNNN
	OPCODE_CXNN
	OPCODE_DXYN
	OPCODE_EX9E
	OPCODE_EXA1
	OPCODE_FX07
	OPCODE_FX0A
	OPCODE_FX15
	OPCODE_FX18
	OPCODE_FX1E
	OPCODE_FX29
	OPCODE_FX33
	OPCODE_FX55
	OPCODE_FX65

	OPCODE_NUMBER
	OPCODE_ERROR
)

var masks = [OPCODE_NUMBER]uint16{0x0000, 0xFFFF, 0xFFFF, 0xF000, 0xF000, 0xF000, 0xF000, 0xF00F, 0xF000, 0xF000, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF00F, 0xF000, 0xF000, 0xF000, 0xF000, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF, 0xF0FF}
var identifiers = [OPCODE_NUMBER]uint16{0x0FFF, 0x00E0, 0x00EE, 0x1000, 0x2000, 0x3000, 0x4000, 0x5000, 0x6000, 0x7000, 0x8000, 0x8001, 0x8002, 0x8003, 0x8004, 0x8005, 0x8006, 0x8007, 0x800E, 0x9000, 0xA000, 0xB000, 0xC000, 0xD000, 0xE09E, 0xE0A1, 0xF007, 0xF00A, 0xF015, 0xF018, 0xF01E, 0xF029, 0xF033, 0xF055, 0xF065}

// Return the opcode associated to the instruction
func GetOpcode(c uint16) (opcode, error) {
	for op, mask := range masks {
		if (mask & c) == identifiers[op] {
			return opcode(op), nil
		}
	}
	return OPCODE_ERROR, fmt.Errorf("unknown opcode %X", c)
}

// Return raw instruction from memory, on 16bits, and increment pointer
func (c *CPU) GetInstruction() uint16 {
	firstByte := uint16(c.Memory[c.Memory_pc]) << 8
	secondByte := uint16(c.Memory[c.Memory_pc+1])

	c.Stack_pc += 2
	return firstByte + secondByte
}
