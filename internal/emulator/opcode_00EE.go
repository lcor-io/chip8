package emulator

// Return from a subroutine
func (e *emulator) opcode_00EE() {

	// Stack is empty, consider this as a NOOP
	if e.Cpu.Stack_pc == 0 {
		return
	}

	e.Cpu.Stack_pc--
	e.Cpu.Memory_pc = e.Cpu.Stack[e.Cpu.Stack_pc] // Move the pc to the old address

}
