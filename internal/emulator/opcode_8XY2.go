package emulator

// Set the register X to register X AND register Y
func (e *emulator) opcode_8XY2(b1 uint8, b2 uint8) {

	e.Cpu.V[b1] = (e.Cpu.V[b1] & e.Cpu.V[b2])

}
