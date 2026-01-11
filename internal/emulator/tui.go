package emulator

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lcor-io/chip8/internal/screen"
)

func (e *emulator) Init() tea.Cmd {
	return nil
}

func (e *emulator) View() string {

	if !e.t_ready {
		return ""
	}

	screenView := e.Screen.View()
	cpuView := e.Cpu.View()
	completeView := lipgloss.JoinHorizontal(lipgloss.Top, screenView, cpuView)

	return lipgloss.Place(e.t_width, e.t_height, 0.5, 0.5, completeView)
}

func (e *emulator) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return e, tea.Quit
		}
	case tea.WindowSizeMsg:
		if !e.t_ready {
			e.t_height = msg.Height
			e.t_width = msg.Width
			e.t_ready = true
		}
	}

	var cmdCpu tea.Cmd
	var cmdScreen tea.Cmd
	var sc = e.Screen.(*screen.Tui)

	_, cmdCpu = e.Cpu.Update(msg)
	_, cmdScreen = sc.Update(msg)
	cmds = append(cmds, cmdCpu, cmdScreen)

	return e, tea.Batch(cmds...)
}
