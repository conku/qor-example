package blogs

import (
	"github.com/conku/page_builder"
	"github.com/conku/publish2"
)

type Page struct {
	page_builder.Page

	publish2.Version
	publish2.Schedule
	publish2.Visible
}
