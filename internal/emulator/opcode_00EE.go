package emulator

// Return from a subroutine
func (e *emulator) opcode_00EE() {

	// Stack is empty, consider this as a NOOP
	if e.Cpu.Sc == 0 {
		return
	}

	e.Cpu.Sc--
	e.Cpu.Pc = e.Cpu.Stack[e.Cpu.Sc] // Move the pc to the old address

}
