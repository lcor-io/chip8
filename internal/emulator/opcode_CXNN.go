package emulator

import "math/rand"

// Set the content of register X to (random byte AND NN)
func (e *emulator) opcode_CXNN(b1 uint8, b2 uint8, b3 uint8) {

	value := (b2 << 4) + b3
	random := uint8(rand.Intn(256))

	e.Cpu.V[b1] = value & random
}
