package table_shortcuts

import (
	"strings"

	"charm.land/lipgloss/v2"
)

func New() string {

	var shortcuts = [][]string{
		{"(r) remove", "(e) edit"},
		{"(c) connect"},
	}

	render := strings.Builder{}

	for i := 0; i < len(shortcuts); i++ {
		for j := 0; j < len(shortcuts[i]); j++ {
			render.WriteString(shortcuts[i][j] + "\t")
		}
		render.WriteString("\n")
	}

	return lipgloss.NewStyle().
		PaddingBottom(2).Render(render.String())

}
