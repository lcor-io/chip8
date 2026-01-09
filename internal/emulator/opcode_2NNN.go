package emulator

import "github.com/lcor-io/chip8/internal/cpu"

// Execute the program at address NNN
func (e *emulator) opcode_2NNN(b1 uint8, b2 uint8, b3 uint8) {

	if e.Cpu.Sc == cpu.STACK_SIZE {
		//TODO Stack overflow, program should panic in some way
		return
	}

	currentAddress := e.Cpu.Pc
	programAddress := (uint16(b1) << 8) + (uint16(b2) << 4) + uint16(b3)

	e.Cpu.Stack[e.Cpu.Sc] = currentAddress // Store the current address in the stack
	e.Cpu.Sc++                             // Increment the stack size
	e.Cpu.Pc = programAddress             // Move the pc to the new address

}
