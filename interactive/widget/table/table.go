package table

import (
	"matman0497/sshboy/interactive/style"

	"charm.land/bubbles/v2/table"
)

func New(tableValues []table.Row) table.Model {

	t := table.New(
		table.WithColumns([]table.Column{{Title: "Server", Width: 30}, {Title: "Host", Width: 30}, {Title: "Index", Width: 0}}),
		table.WithRows(tableValues),
		table.WithFocused(true),
		table.WithWidth(60),
	)

	s := table.DefaultStyles()
	s.Selected = s.Selected.
		Foreground(style.PrimaryColor).
		Background(style.TableBackground).
		Bold(false)

	t.SetStyles(s)

	return t

}
