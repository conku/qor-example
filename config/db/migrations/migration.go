package migrations

import (
	"github.com/conku/activity"
	"github.com/conku/auth/auth_identity"
	"github.com/conku/banner_editor"
	"github.com/conku/help"
	"github.com/conku/media/asset_manager"

	"github.com/conku/qor-example/app/admin"
	"github.com/conku/qor-example/config/db"
	"github.com/conku/qor-example/models/blogs"
	"github.com/conku/qor-example/models/orders"
	"github.com/conku/qor-example/models/products"
	"github.com/conku/qor-example/models/seo"
	"github.com/conku/qor-example/models/settings"
	"github.com/conku/qor-example/models/stores"
	"github.com/conku/qor-example/models/users"
	"github.com/conku/transition"
)

func init() {
	AutoMigrate(&asset_manager.AssetManager{})

	AutoMigrate(&products.Product{}, &products.ProductVariation{}, &products.ProductImage{}, &products.ColorVariation{}, &products.ColorVariationImage{}, &products.SizeVariation{})
	AutoMigrate(&products.Color{}, &products.Size{}, &products.Material{}, &products.Category{}, &products.Collection{})

	AutoMigrate(&users.User{}, &users.Address{})
	AutoMigrate(&orders.Order{}, &orders.OrderItem{})
	AutoMigrate(&orders.DeliveryMethod{})
	AutoMigrate(&stores.Store{})
	AutoMigrate(&settings.Setting{}, &settings.MediaLibrary{})
	AutoMigrate(&transition.StateChangeLog{})
	AutoMigrate(&activity.QorActivity{})

	AutoMigrate(&admin.QorWidgetSetting{})

	AutoMigrate(&blogs.Page{}, &blogs.Article{})

	AutoMigrate(&seo.MySEOSetting{})

	AutoMigrate(&help.QorHelpEntry{})

	AutoMigrate(&auth_identity.AuthIdentity{})

	AutoMigrate(&banner_editor.QorBannerEditorSetting{})
}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
