package emulator

// Add the content of the register X to the content of the register Y, and set the result to register X. Set carry register to 1 if result is > 8bits
func (e *emulator) opcode_8XY4(b1 uint8, b2 uint8) {

	sum := uint16(e.Cpu.V[b1]) + uint16(e.Cpu.V[b2])

	if sum > 0xFF {
		e.Cpu.V[0xF] = 1
	}

	e.Cpu.V[b1] = uint8(sum)

}
