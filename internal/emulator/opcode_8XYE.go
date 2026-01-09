package emulator

// If the most-significant bit of register X is 1, carry register is set to 1. The content of register X is then multiplied by 2
func (e *emulator) opcode_8XYE(b1 uint8) {

	e.Cpu.V[0xF] = e.Cpu.V[b1] >> 7
	e.Cpu.V[b1] = e.Cpu.V[b1] / 2

}
