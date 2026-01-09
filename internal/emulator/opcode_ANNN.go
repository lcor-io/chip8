package emulator

// Set the address register to the value of NNN
func (e *emulator) opcode_ANNN(b1 uint8, b2 uint8, b3 uint8) {
	newAddress := (uint16(b1) << 8) + (uint16(b2) << 4) + uint16(b3)
	e.Cpu.I = newAddress
}
