package admin

import (
	"fmt"

	"github.com/qor/admin"
	"github.com/qor/media_library"
	"github.com/qor/qor"
	"github.com/qor/qor-example/app/models"
	"github.com/qor/qor-example/db"
	"github.com/qor/widget"
	"html/template"
)

var Widgets *widget.Widgets

func init() {
	Widgets = widget.New(&widget.Config{DB: db.DB})
	Admin.AddResource(Widgets)

	// Top Banner
	type bannerArgument struct {
		Title           string
		ButtonTitle     string
		Link            string
		BackgroundImage media_library.FileSystem
		Logo            media_library.FileSystem
	}

	Widgets.RegisterWidget(&widget.Widget{
		Name:     "Banner",
		Template: "banner",
		Setting:  Admin.NewResource(&bannerArgument{}),
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			context.Options["Setting"] = setting
			return context
		},
	})

	// Banner Editor
	type bannerEditorArgument struct {
		Value string
	}
	bannerEditorResource := Admin.NewResource(&bannerEditorArgument{})
	bannerEditorResource.Meta(&admin.Meta{Name: "Value", Type: "banner_editor"})

	Widgets.RegisterWidget(&widget.Widget{
		Name:     "BannerEditor",
		Template: "banner_editor",
		Setting:  bannerEditorResource,
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			context.Options["Value"] = template.HTML(setting.(*bannerEditorArgument).Value)
			return context
		},
	})

	// selected Products
	type selectedProductsArgument struct {
		Products []string
	}
	selectedProductsResource := Admin.NewResource(&selectedProductsArgument{})
	selectedProductsResource.Meta(&admin.Meta{Name: "Products", Type: "select_many", Collection: func(value interface{}, context *qor.Context) [][]string {
		var collectionValues [][]string
		var products []*models.Product
		db.DB.Find(&products)
		for _, product := range products {
			collectionValues = append(collectionValues, []string{fmt.Sprintf("%v", product.ID), product.Name})
		}
		return collectionValues
	}})
	Widgets.RegisterWidget(&widget.Widget{
		Name:     "Products",
		Template: "products",
		Setting:  selectedProductsResource,
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			if setting != nil {
				var products []*models.Product
				db.DB.Limit(9).Preload("ColorVariations").Preload("ColorVariations.Images").Where("id IN (?)", setting.(*selectedProductsArgument).Products).Find(&products)
				context.Options["Products"] = products
			}
			return context
		},
	})
}
