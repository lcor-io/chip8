package emulator

// Place the value of the Delay Timer in register X
func (e *emulator) opcode_FX07(b1 uint8) {

	e.Cpu.V[b1] = e.Cpu.Delay_timer

}
