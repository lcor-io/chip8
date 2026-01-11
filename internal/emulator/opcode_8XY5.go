package emulator

// Substract the content of the register Y to the content of the register X, and set the result to register X. Set carry register to 1 if content of register X > content of register Y
func (e *emulator) opcode_8XY5(b1 uint8, b2 uint8) {

	if e.Cpu.V[b1] > e.Cpu.V[b2] {
		e.Cpu.V[0xF] = 1
	} else {
		e.Cpu.V[0xF] = 0
	}

	e.Cpu.V[b1] = e.Cpu.V[b1] - e.Cpu.V[b2]

}
