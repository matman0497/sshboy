package interactive

import (
	"matman0497/sshboy/internal"
	"matman0497/sshboy/internal/config"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

func HandleServerListUpdate(m Model, msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "c":
			//connect to server via ssh
			server := config.GetServer(m.serverListModel.table.SelectedRow()[0])

			clearTerminal()

			return m, tea.ExecProcess(internal.Connect(server), func(err error) tea.Msg {
				return SshConnectionFinished{err: err}
			})

		case "e":
			//edit the server
			m.currentState = EditServer

			server := config.GetServer(m.serverListModel.table.SelectedRow()[0])

			m.serverEditModel.serverInEdit = server

			//set input value to current value from the config
			for i := range m.serverEditModel.inputs {

				switch i {
				case 0:
					input := m.serverEditModel.inputs[i]
					input.SetValue(server.Name)
					m.serverEditModel.inputs[i] = input
				case 1:
					input := m.serverEditModel.inputs[i]
					input.SetValue(server.Host)
					m.serverEditModel.inputs[i] = input
				}

			}

			return m, nil

		case "r":
			//delete the server from the config
			var serverNameToDelete = m.serverListModel.table.SelectedRow()[0]
			config.DeleteServer(serverNameToDelete)
			config.Save()

			//update the table view by removing the deleted row
			//TODO: consider if it would be better/easier to just reload all servers from the config
			var newRows []table.Row
			for _, row := range m.serverListModel.table.Rows() {
				if row[0] != serverNameToDelete {
					newRows = append(newRows, row)
				}
			}

			m.serverListModel.table.SetRows(newRows)

		}
	}

	return m, nil
}
