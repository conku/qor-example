package i18n

import (
	"path/filepath"

	"github.com/conku/i18n"
	"github.com/conku/i18n/backends/database"
	"github.com/conku/i18n/backends/yaml"

	"github.com/conku/qor-example/config"
	"github.com/conku/qor-example/config/db"
)

var I18n *i18n.I18n

func init() {
	I18n = i18n.New(database.New(db.DB), yaml.New(filepath.Join(config.Root, "config/locales")))
}
