package interactive

import (
	"fmt"
	"mattiamancina/sshboy/internal/config"
	"os"
	"sort"
	"strings"

	"charm.land/bubbles/v2/table"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/common-nighthawk/go-figure"
)

var (
	focusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle         = focusedStyle
	noStyle             = lipgloss.NewStyle()
	helpStyle           = blurredStyle
	cursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type State int

const (
	Start      State = 0
	Index      State = 1
	EditServer State = 2
)

var stateName = map[State]string{
	Start:      "start",
	Index:      "index",
	EditServer: "edit-server",
}

func initialModel(t table.Model, inputs []textinput.Model) Model {
	return Model{
		serverListModel: ServerListModel{
			table: t,
		},
		serverEditModel: ServerEditModel{
			inputs: inputs,
		},
		currentState: Start,
	}
}

type SshConnectionFinished struct{ err error }

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	//check quit first
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		}
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
		m, cmd := HandleServerListUpdate(m, msg)

		//if the handler returns a command return it, otherwise update the table
		if cmd != nil {
			return m, cmd
		}

		m.serverListModel.table, cmd = m.serverListModel.table.Update(msg)

		return m, nil

	case EditServer:
		//the edit server view
		m, cmd := HandleServerEditUpdate(m, msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() tea.View {

	logo := figure.NewFigure("sshboy", "", true)
	switch m.currentState {
	case Start:

	case Index:
		return tea.NewView(logo.String() + m.serverListModel.table.View() + m.serverListModel.table.HelpView())

	case EditServer:
		var b strings.Builder
		var c *tea.Cursor

		for i, in := range m.serverEditModel.inputs {
			b.WriteString(m.serverEditModel.inputs[i].View())
			if i < len(m.serverEditModel.inputs)-1 {
				b.WriteRune('\n')
			}

			if in.Focused() {
				c = in.Cursor()
				if c != nil {
					c.Y += i
				}
			}
		}

		button := &blurredButton
		if m.serverEditModel.focusIndex == len(m.serverEditModel.inputs) {
			button = &focusedButton
		}

		fmt.Fprintf(&b, "\n\n%s\n\n", *button)

		v := tea.NewView(b.String())
		v.Cursor = c
		return v
	}

	return tea.NewView("whoops!")
}

func Init() {

	//init the table
	var tableValues []table.Row

	var servers = config.Get().Servers

	sort.Slice(servers, func(i, j int) bool {
		return config.Get().Servers[i].Name < config.Get().Servers[j].Name
	})
	for i, server := range servers {
		tableValues = append(tableValues, table.Row{server.Name, server.Host, string(rune(i))})
	}

	t := table.New(
		table.WithColumns([]table.Column{{Title: "Server", Width: 30}, {Title: "Host", Width: 30}, {Title: "Index", Width: 0}}),
		table.WithRows(tableValues),
		table.WithFocused(true),
		table.WithWidth(60),
	)
	s := table.DefaultStyles()
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.BrightBlue).
		Bold(false)

	t.SetStyles(s)

	//init the input fields for the edit view
	var inputs []textinput.Model

	//server name input
	var nameInput textinput.Model
	nameInput = textinput.New()
	nameInput.Prompt = "Name: "
	nameInput.CharLimit = 32

	textinputStyle := nameInput.Styles()
	textinputStyle.Cursor.Color = lipgloss.Color("205")
	textinputStyle.Focused.Prompt = focusedStyle
	textinputStyle.Focused.Text = focusedStyle
	textinputStyle.Blurred.Prompt = blurredStyle
	textinputStyle.Focused.Text = focusedStyle
	nameInput.SetStyles(textinputStyle)

	//server host input
	var hostInput textinput.Model
	hostInput = textinput.New()
	hostInput.Prompt = "Host: "
	hostInput.CharLimit = 32

	hostInputStyle := hostInput.Styles()
	hostInputStyle.Cursor.Color = lipgloss.Color("205")
	hostInputStyle.Focused.Prompt = focusedStyle
	hostInputStyle.Focused.Text = focusedStyle
	hostInputStyle.Blurred.Prompt = blurredStyle
	hostInputStyle.Focused.Text = focusedStyle
	hostInput.SetStyles(hostInputStyle)

	inputs = append(inputs, nameInput)
	inputs = append(inputs, hostInput)

	p := tea.NewProgram(initialModel(t, inputs))

	if _, err := p.Run(); err != nil {
		fmt.Printf("There's been an error: %v", err)
		os.Exit(1)
	}
}

func clearTerminal() {
	fmt.Print("\033[H\033[2J")
}
