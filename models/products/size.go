package products

import (
	"strings"

	"github.com/conku/l10n"
	"github.com/conku/sorting"
	"github.com/conku/validations"
	"github.com/jinzhu/gorm"
)

type Size struct {
	gorm.Model
	l10n.Locale
	sorting.Sorting
	Name string
	Code string `l10n:"sync"`
}

func (size Size) Validate(db *gorm.DB) {
	if strings.TrimSpace(size.Name) == "" {
		db.AddError(validations.NewError(size, "Name", "Name can not be empty"))
	}

	if strings.TrimSpace(size.Code) == "" {
		db.AddError(validations.NewError(size, "Code", "Code can not be empty"))
	}
}
