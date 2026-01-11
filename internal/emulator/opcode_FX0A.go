package emulator

import "time"

// Wait for a key press, and store the value of the key in Vx
func (e *emulator) opcode_FX0A(b1 uint8) {

	oldKeyboard := e.Keyboard.Keys
	go func() {
		keyboardEquals := true
		e.Stop()

		for keyboardEquals {
			for i, key := range e.Keyboard.Keys {
				if oldKeyboard[i] != key {
					e.Cpu.V[b1] = uint8(i)
					keyboardEquals = false
				}
			}
			time.Sleep(time.Millisecond * 5)
		}
		e.Start()
	}()

}
