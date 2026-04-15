package interactive

import (
	"fmt"
	"log"
	"matman0497/sshboy/interactive/widget/button"
	"matman0497/sshboy/interactive/widget/input"
	"matman0497/sshboy/internal/config"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

type ServerEditModel struct {
	server     *config.Server
	inputs     []Input
	focusIndex int
}

type Input struct {
	Key   string
	Input *textinput.Model
}

func newServerEditModel() ServerEditModel {
	nameInput := input.New("Name: ")
	nameInput.Focus()

	hostInput := input.New("Host: ")

	return ServerEditModel{
		inputs: []Input{
			{Key: "nameInput", Input: &nameInput},
			{Key: "hostnameInput", Input: &hostInput},
		},
		focusIndex: 0,
	}
}

func (m ServerEditModel) View() tea.View {

	var b strings.Builder
	var c *tea.Cursor

	for key, inp := range m.inputs {

		b.WriteString(m.inputs[key].Input.View())

		if key < len(m.inputs)-1 {
			b.WriteRune('\n')
		}

		if inp.Input.Focused() {
			c = inp.Input.Cursor()
			if c != nil {
				c.Y += key
			}
		}

	}

	bt := button.NewBlurred("Submit")
	if m.focusIndex == len(m.inputs) {
		bt = button.NewFocused("Submit")
	}

	fmt.Fprintf(&b, "\n\n%s\n\n", bt)

	v := tea.NewView(b.String())
	v.Cursor = c
	return v
}

func (m ServerEditModel) Update(msg tea.Msg) (ServerEditModel, tea.Cmd) {
	log.Println("1")
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "esc":

			return m, tea.Quit

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, save.
			if s == "enter" && m.focusIndex == len(m.inputs) {

				//return save message here
				return m, func() tea.Msg {
					return SaveRequested{
						server: m.server,
						//TODO: maybe introduce get input by value
						host: m.inputs[0].Input.Value(),
						name: m.inputs[1].Input.Value(),
					}
				}
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))

			for key, _ := range m.inputs {

				if key == m.focusIndex {
					// Set focused state
					cmds[key] = m.inputs[key].Input.Focus()
					continue
				}

				m.inputs[key].Input.Blur()

			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m ServerEditModel) Init() tea.Cmd {

	return nil
}

func (m *ServerEditModel) updateInputs(msg tea.Msg) tea.Cmd {

	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		*m.inputs[i].Input, cmds[i] = m.inputs[i].Input.Update(msg)

	}

	return tea.Batch(cmds...)
}
