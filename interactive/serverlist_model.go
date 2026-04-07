package interactive

import (
	"mattiamancina/sshboy/internal/config"

	"charm.land/bubbles/v2/table"
)

type ServerListModel struct {
	config   *config.Config
	selected map[int]struct{}
	table    table.Model
}
