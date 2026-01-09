package emulator

// Set Delay timer to the content of register X
func (e *emulator) opcode_FX15(b1 uint8) {

	e.Cpu.Delay_timer = e.Cpu.V[b1]

}
