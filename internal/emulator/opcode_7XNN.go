package emulator

// Add the constant NN to the content of the register X
func (e *emulator) opcode_7XNN(b1 uint8, b2 uint8, b3 uint8) {

	value := (b2 << 4) + b3
	e.Cpu.V[b1] = e.Cpu.V[b1] + value

}
