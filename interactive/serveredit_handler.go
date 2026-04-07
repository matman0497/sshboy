package interactive

import (
	"matman0497/sshboy/internal/config"

	tea "charm.land/bubbletea/v2"
)

func HandleServerEditUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
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
			if s == "enter" && m.serverEditModel.focusIndex == len(m.serverEditModel.inputs) {

				save(m)
				m.currentState = Index
				return m, nil
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.serverEditModel.focusIndex--
			} else {
				m.serverEditModel.focusIndex++
			}

			if m.serverEditModel.focusIndex > len(m.serverEditModel.inputs) {
				m.serverEditModel.focusIndex = 0
			} else if m.serverEditModel.focusIndex < 0 {
				m.serverEditModel.focusIndex = len(m.serverEditModel.inputs)
			}

			cmds := make([]tea.Cmd, len(m.serverEditModel.inputs))
			for i := 0; i <= len(m.serverEditModel.inputs)-1; i++ {
				if i == m.serverEditModel.focusIndex {
					// Set focused state
					cmds[i] = m.serverEditModel.inputs[i].Focus()
					continue
				}
				// Remove focused state
				m.serverEditModel.inputs[i].Blur()
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func save(m Model) {
	for i, input := range m.serverEditModel.inputs {
		switch i {
		case 0:
			m.serverEditModel.serverInEdit.SetName(input.Value())

			config.Save()
		case 1:
			m.serverEditModel.serverInEdit.SetHost(input.Value())

			config.Save()
		}
	}
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.serverEditModel.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.serverEditModel.inputs {
		m.serverEditModel.inputs[i], cmds[i] = m.serverEditModel.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
