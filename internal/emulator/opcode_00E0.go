package emulator

// Clear the screen
func (e *emulator) opcode_00E0() {

	e.Screen.Clear()

}
