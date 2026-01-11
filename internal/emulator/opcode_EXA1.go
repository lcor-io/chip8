package emulator

// Skip next instruction if key with value at register X is not pressed
func (e *emulator) opcode_EXA1(b1 uint8) {

	key := e.Cpu.V[b1]
	if !e.Keyboard.Keys[key] {
		e.Cpu.Pc += 2
	}

}
