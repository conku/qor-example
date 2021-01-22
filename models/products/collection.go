package products

import (
	"github.com/conku/gorm"
	"github.com/conku/l10n"
)

type Collection struct {
	gorm.Model
	Name string
	l10n.LocaleCreatable
}
