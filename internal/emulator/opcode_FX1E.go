package emulator

// Add the content of the register address register and register X, and set the result to the address register
func (e *emulator) opcode_FX1E(b1 uint8) {

	e.Cpu.I = uint16(e.Cpu.V[b1]) + e.Cpu.I

}
