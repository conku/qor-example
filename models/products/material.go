package products

import (
	"github.com/conku/gorm"
	"github.com/conku/l10n"
)

type Material struct {
	gorm.Model
	l10n.Locale
	Name string
	Code string `l10n:"sync"`
}
