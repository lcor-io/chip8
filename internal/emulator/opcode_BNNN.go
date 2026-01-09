package emulator

// Move the program counter at the address NNN + register 0
func (e *emulator) opcode_BNNN(b1 uint8, b2 uint8, b3 uint8) {
	newAddress := (uint16(b1) << 8) + (uint16(b2) << 4) + uint16(b3)
	e.Cpu.Pc = newAddress + uint16(e.Cpu.V[0])
}
