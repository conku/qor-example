package blogs

import (
	"github.com/conku/gorm"
	"github.com/conku/publish2"
	"github.com/conku/qor-example/models/users"
)

type Article struct {
	gorm.Model
	Author   users.User
	AuthorID uint
	Title    string
	Content  string `gorm:"type:text"`
	publish2.Version
	publish2.Schedule
	publish2.Visible
}
