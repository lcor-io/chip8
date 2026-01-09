package emulator

// The value of the address register is set to the location for the hexadecimal sprite corresponding to the value of register X
func (e *emulator) opcode_FX29(b1 uint8) {

	e.Cpu.I = 5 * uint16(e.Cpu.V[b1])
}
