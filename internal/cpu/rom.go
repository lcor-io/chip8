package cpu

import "os"

// Load a rom in the internal memory
func (c *CPU) LoadRom(p string) error {

	file, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Read(c.Memory[INITIAL_RAM_ADDRESS:])
	if err != nil {
		panic(err)
	}

	return nil
}
