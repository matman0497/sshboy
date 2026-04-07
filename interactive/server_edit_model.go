package interactive

import (
	"matman0497/sshboy/internal/config"

	"charm.land/bubbles/v2/textinput"
)

type ServerEditModel struct {
	server       *config.Server
	inputs       []textinput.Model
	focusIndex   int
	serverInEdit *config.Server
}
