package style

import (
	"charm.land/lipgloss/v2"
)

var (
	PrimaryColor = lipgloss.Color("39")
	FocusedStyle = lipgloss.NewStyle().Foreground(PrimaryColor)
	BlurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))

	TableBackground = lipgloss.Color("240")
)
