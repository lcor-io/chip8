package emulator

// Store the content of register 0 to X in memory, starting from address I
func (e *emulator) opcode_FX55(b1 uint8) {

	for reg := range b1 {
		e.Cpu.Memory[e.Cpu.I+uint16(reg)] = e.Cpu.V[reg]
	}
	e.Cpu.Memory[e.Cpu.I+uint16(b1)] = e.Cpu.V[b1]

}
