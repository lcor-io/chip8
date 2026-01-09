package emulator

// Set Sound timer to the content of register X
func (e *emulator) opcode_FX18(b1 uint8) {

	e.Cpu.Sound_timer = e.Cpu.V[b1]

}
