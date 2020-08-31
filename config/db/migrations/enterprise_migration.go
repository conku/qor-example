// +build enterprise

package migrations

import "github.com/conku/qor-example/app/enterprise"

func init() {
	AutoMigrate(&enterprise.QorMicroSite{})
}
