package interactive

import (
	tableComponent "matman0497/sshboy/interactive/widget/table"
	"matman0497/sshboy/internal"
	"matman0497/sshboy/internal/config"
	"sort"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
)

type ServerListModel struct {
	config   *config.Config
	selected map[int]struct{}
	table    table.Model
	store    config.ServerStore
}

func newServerListModel() ServerListModel {

	//init the table
	var tableValues []table.Row

	store := config.Store{}

	var servers = store.List()

	sort.Slice(servers, func(i, j int) bool {
		return store.List()[i].Name < store.List()[j].Name
	})

	for i, server := range servers {
		tableValues = append(tableValues, table.Row{server.Name, server.Host, string(rune(i))})
	}

	t := tableComponent.New(tableValues)

	return ServerListModel{
		table: t,
		store: store,
	}
}

func (m ServerListModel) Update(msg tea.Msg) (ServerListModel, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {

		case "ctrl+c", "esc":

			return m, tea.Quit

		case "c":
			//connect to server via ssh
			server := config.GetServer(m.table.SelectedRow()[0])

			clearTerminal()

			return m, tea.ExecProcess(internal.Connect(server), func(err error) tea.Msg {
				return SshConnectionFinished{err: err}
			})

		case "e":
			server := config.GetServer(m.table.SelectedRow()[0])
			return m, func() tea.Msg {
				return EditRequested{
					server: server,
				}
			}

		case "r":
			//delete the server from the config
			server := config.GetServer(m.table.SelectedRow()[0])
			return m, func() tea.Msg {
				return DeleteRequested{
					server: server,
				}
			}

		}

		m.table, _ = m.table.Update(msg)
	}

	return m, nil
}

func (m ServerListModel) View() tea.View {
	return tea.NewView(m.table.View())

}

func (m ServerListModel) Init() tea.Cmd {
	return nil
}

func (m ServerListModel) Refresh() ServerListModel {

	var servers = m.store.List()
	var tableValues []table.Row

	sort.Slice(servers, func(i, j int) bool {
		return m.store.List()[i].Name < m.store.List()[j].Name
	})

	for i, server := range servers {
		tableValues = append(tableValues, table.Row{server.Name, server.Host, string(rune(i))})
	}

	m.table.SetRows(tableValues)

	return m
}
