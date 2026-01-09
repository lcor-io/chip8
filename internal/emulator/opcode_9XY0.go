package emulator

// Jump next instruction if VX != VY
func (e *emulator) opcode_9XY0(b1 uint8, b2 uint8) {

	registerXValue := e.Cpu.V[b1]
	registerYValue := e.Cpu.V[b2]

	if registerXValue != registerYValue {
		e.Cpu.Pc += 2
	}

}
