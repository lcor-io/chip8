package emulator

// Skip next instruction if key with value at register X is pressed
func (e *emulator) opcode_EX9E(b1 uint8) {

	key := e.Cpu.V[b1]
	if e.Keyboard.Keys[key] {
		e.Cpu.Pc += 2
	}

}
