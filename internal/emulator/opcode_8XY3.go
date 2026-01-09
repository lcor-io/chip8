package emulator

// Set the register X to register X XOR register Y
func (e *emulator) opcode_8XY3(b1 uint8, b2 uint8) {

	e.Cpu.V[b1] = (e.Cpu.V[b1] ^ e.Cpu.V[b2])

}
