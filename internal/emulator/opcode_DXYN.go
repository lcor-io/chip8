package emulator

const SPRITE_LENGTH uint8 = 8

func (e *emulator) opcode_DXYN(x uint8, y uint8, spriteHeight uint8) {

	e.Cpu.V[0xF] = 0 // Set the carry-over register to 0

	for row := range spriteHeight {

		spriteRow := e.Cpu.Memory[e.Cpu.I+uint16(row)] // Loop on each line of the sprite in memory
		py_coordinate := e.Cpu.V[y] + row

		for column := range SPRITE_LENGTH { // Loop on each column of the sprite at the current row

			px_coordinate := e.Cpu.V[x] + column
			current_sprite_value := ((spriteRow << column) & 0b10000000) // Get the bit of the row at the column position

			if current_sprite_value == 0 { // The sprite is not lit, we continue the iteration
				continue
			}

			// The sprite is lit, we should update the screen

			// The pixel is already set, increment the carry-over register
			current_screen_value := e.Screen.GetPixel(int(px_coordinate), int(py_coordinate))
			if current_screen_value {
				e.Cpu.V[0xF] = 1
			}

			e.Screen.SetPixel(int(px_coordinate), int(py_coordinate), !current_screen_value)
		}

	}

}
