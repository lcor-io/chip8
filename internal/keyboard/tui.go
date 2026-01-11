package keyboard

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lcor-io/chip8/internal/utils/styles"
)

func (k *Keyboard) Init() tea.Cmd {
	return nil
}

func (k *Keyboard) View() string {

	keyOrder := [][]int{{1, 2, 3, 12}, {4, 5, 6, 13}, {7, 8, 9, 14}, {10, 0, 11, 15}}

	rows := []string{}
	for _, row := range keyOrder {

		cols := []string{}

		for _, col := range row {
			isKeyPressed := k.Keys[col]
			if isKeyPressed {
				cols = append(cols, styles.KeyboardPressedCell.Render(fmt.Sprintf("%X", col)))
			} else {
				cols = append(cols, styles.KeyboardCell.Render(fmt.Sprintf("%X", col)))
			}
		}

		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, cols...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
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
			k.keyDebounce[key] = 60
		}
	}
	return k, nil
}
