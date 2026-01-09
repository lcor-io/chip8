package emulator

// Set the content of register 0 to X with the values in memory stored from address I
func (e *emulator) opcode_FX65(b1 uint8) {

	for reg := range b1 {
		e.Cpu.V[reg] = e.Cpu.Memory[e.Cpu.I+uint16(reg)]
	}
	e.Cpu.V[b1] = e.Cpu.Memory[e.Cpu.I+uint16(b1)]
}
