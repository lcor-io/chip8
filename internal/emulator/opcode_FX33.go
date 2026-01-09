package emulator

// Store BCD representation of the value of register X in memory location I,I+1,I+2
func (e *emulator) opcode_FX33(b1 uint8) {

	val := e.Cpu.V[b1]
	e.Cpu.Memory[e.Cpu.I] = val / 100         // hundreds
	e.Cpu.Memory[e.Cpu.I+1] = (val / 10) % 10 // tens
	e.Cpu.Memory[e.Cpu.I+2] = val % 10        // ones

}
