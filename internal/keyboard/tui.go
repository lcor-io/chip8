package keyboard

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

func (k *Keyboard) Init() tea.Cmd {
	return nil
}

func (k *Keyboard) View() string {
	return ""
}

func (k *Keyboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "a", "b", "c", "d", "e", "f":
			key, _ := strconv.ParseInt(msg.String(), 16, 8) // Convert keypress to it's hexadecimal value
			if key > 16 {
				return k, nil
			}
			k.Keys[key] = true // Set the key as pressed on the keyboard
		}
	}
	return k, nil
}
