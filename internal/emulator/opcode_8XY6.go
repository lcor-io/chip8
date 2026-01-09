package emulator

// If the least-significant bit of register X is 1, carry register is set to 1. The content of register X is then divided by 2
func (e *emulator) opcode_8XY6(b1 uint8) {

	e.Cpu.V[0xF] = e.Cpu.V[b1] & 0b00000001
	e.Cpu.V[b1] = e.Cpu.V[b1] / 2

}
