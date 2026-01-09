package emulator

// Jump next instruction if VX != NN
func (e *emulator) opcode_4XNN(b1 uint8, b2 uint8, b3 uint8) {

	registerValue := e.Cpu.V[b1]
	value := ((uint8(b2) << 4) + uint8(b3))

	if registerValue != value {
		e.Cpu.Memory_pc += 2
	}

}
