package orders

import (
	"github.com/conku/gorm"
)

type DeliveryMethod struct {
	gorm.Model

	Name  string
	Price float32
}
