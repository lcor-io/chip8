package emulator

// Set the value NN to register X
func (e *emulator) opcode_6XNN(b1 uint8, b2 uint8, b3 uint8) {

	value := ((uint8(b2) << 4) + uint8(b3))
	e.Cpu.V[b1] = value

}
