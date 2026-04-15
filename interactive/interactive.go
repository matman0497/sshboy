package interactive

import (
	"fmt"
	"matman0497/sshboy/interactive/style"
	"matman0497/sshboy/interactive/widget/table_shortcuts"
	"matman0497/sshboy/internal/config"
	"os"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/common-nighthawk/go-figure"
)

type State int

const (
	Start      State = 0
	Index      State = 1
	EditServer State = 2
)

func initialModel() Model {
	return Model{
		serverListModel: newServerListModel(),
		serverEditModel: newServerEditModel(),
		currentState:    Start,
	}
}

type SshConnectionFinished struct{ err error }
type EditRequested struct {
	server *config.Server
}

type DeleteRequested struct {
	server *config.Server
}

type SaveRequested struct {
	server *config.Server
	name   string
	host   string
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		}

	case EditRequested:
		m.currentState = EditServer

		m.serverEditModel.server = msg.server
		m.serverEditModel.inputs[0].Input.SetValue(msg.server.Host)
		m.serverEditModel.inputs[1].Input.SetValue(msg.server.Name)

		return m, nil

	case SaveRequested:

		m.currentState = Index

		msg.server.Name = msg.name
		msg.server.Host = msg.host

		m.serverListModel.store.Save()

		m.serverListModel = m.serverListModel.Refresh()
		return m, nil

	case DeleteRequested:

		m.currentState = Index

		m.serverListModel.store.Delete(msg.server.Name)
		m.serverListModel.store.Save()

		m.serverListModel = m.serverListModel.Refresh()
		return m, nil

	}

	//check by state
	switch m.currentState {
	case Start:
		//the application has been started. Set the current start to index of servers which is the start view
		m.currentState = Index
		clearTerminal()
		return m, tea.ClearScreen
	case Index:
		//the index view. Here we display a table of all the servers found in the config
		var cmd tea.Cmd
		m.serverListModel, cmd = m.serverListModel.Update(msg)

		return m, cmd

	case EditServer:
		//the index view. Here we display a table of all the servers found in the config
		var cmd tea.Cmd
		m.serverEditModel, cmd = m.serverEditModel.Update(msg)

		return m, cmd
	}

	return m, nil
}

func (m Model) View() tea.View {

	logo := figure.NewFigure("sshboy", "", true)

	switch m.currentState {
	case Start:

	case Index:
		return tea.NewView(lipgloss.NewStyle().Foreground(style.PrimaryColor).Render(logo.String()) + "\n" +
			table_shortcuts.New() + "\n" + m.serverListModel.View().Content)

	case EditServer:
		return m.serverEditModel.View()
	}

	return tea.NewView("whoops!")
}

func Init() {

	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
