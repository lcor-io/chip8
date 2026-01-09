package cpu

import "testing"

func TestLoadingRom(t *testing.T) {

	cpu := New()

	t.Run("Load rom in memory", func(t *testing.T) {
		t.Chdir("../../roms")
		err := cpu.LoadRom("maze.ch8")
		if err != nil {
			t.Errorf("Error while loading rom in memory: %v", err)
		}
	})

	t.Run("Initial address is correct", func(t *testing.T) {
		if cpu.Pc != INITIAL_RAM_ADDRESS {
			t.Errorf("Incorrect address, [%d] wanted, [%d] found", INITIAL_RAM_ADDRESS, cpu.Pc)
		}
	})

	t.Run("First bytes are empty", func(t *testing.T) {
		for i := range cpu.Memory[:INITIAL_RAM_ADDRESS] {
			if cpu.Memory[i] != 0x0 {
				t.Errorf("Expecting empty byte at address %d", i)
			}
		}
	})

	t.Run("Rom size is correct", func(t *testing.T) {
		const ROM_SIZE uint16 = 38
		const LAST_INSTRUCTION_ADDRESS = INITIAL_RAM_ADDRESS + ROM_SIZE - 1

		firstInstruction := cpu.Memory[INITIAL_RAM_ADDRESS]
		lastInstruction := cpu.Memory[LAST_INSTRUCTION_ADDRESS]
		afterLastInstruction := cpu.Memory[LAST_INSTRUCTION_ADDRESS+1]
		if firstInstruction == 0x0 || lastInstruction == 0x0 || afterLastInstruction != 0x0 {
			t.Error("Rom size is not correct")
		}
	})
}
