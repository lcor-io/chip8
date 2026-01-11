package cpu

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
	"github.com/lcor-io/chip8/internal/screen"
	"github.com/lcor-io/chip8/internal/utils/styles"
)

const (
	address_to_show = 6
)

func (c *CPU) Init() tea.Cmd {
	return nil
}

func (c *CPU) View() string {

	if !c.ready {
		return ""
	}

	memoryView := ""
	for i, byte := range c.Memory {
		formatted := fmt.Sprintf("%.2X", byte)
		if i == int(c.Pc) {
			formatted = styles.PCMemoryCell.Render(formatted)
		}
		memoryView = memoryView + " " + formatted

		if i != 0 && (i+1)%address_to_show == 0 {
			memoryView = memoryView + "\n"
		}
	}
	c.viewport.SetContent(memoryView)
	memoryView = styles.Border.Render(c.viewport.View())

	registerList := list.New().
		Enumerator(func(l list.Items, i int) string {
			return fmt.Sprintf("V%X ->", i)
		}).
		EnumeratorStyle(styles.RegisterListEnumerator).
		ItemStyle(styles.RegisterListItem)
	for _, register := range c.V {
		registerList.Item(fmt.Sprintf("0x%.2X", register))
	}

	signalList := list.New(
		styles.RegisterListEnumerator.Render("Delay ->")+styles.RegisterListItem.Render(fmt.Sprintf("%d", c.Delay_timer)),
		styles.RegisterListEnumerator.Render("Sound ->")+styles.RegisterListItem.Render(fmt.Sprintf("%d", c.Sound_timer))).
		Enumerator(func(items list.Items, index int) string { return "" }).
		EnumeratorStyle(styles.RegisterListEnumerator)

	registerView := lipgloss.JoinVertical(lipgloss.Left,
		styles.RegisterTitle.Render("Registers"),
		registerList.String(),
		styles.RegisterTitle.Render("Address register"),
		fmt.Sprintf("%s0x%.4X", styles.RegisterListEnumerator.Render("I ->"), c.I),
		styles.RegisterTitle.Render("Program counter"),
		fmt.Sprintf("%s%.4d", styles.RegisterListEnumerator.Render("PC ->"), c.Pc),
		styles.RegisterTitle.Render("Signals"),
		signalList.String(),
	)

	return lipgloss.JoinHorizontal(lipgloss.Top, memoryView, registerView)
}

func (c *CPU) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg.(type) {
	case tea.WindowSizeMsg:
		if !c.ready {
			c.viewport = viewport.New((address_to_show*3)+1, int(screen.SCREEN_HEIGHT))
			c.ready = true
		} else {
			c.viewport.Width = (address_to_show * 3) + 1
			c.viewport.Height = int(screen.SCREEN_HEIGHT)
		}
	}

	var cmd tea.Cmd
	c.viewport, cmd = c.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return c, tea.Batch(cmds...)
}
