package auth

import (
	"time"

	"github.com/conku/auth"
	"github.com/conku/auth/authority"
	"github.com/conku/auth/providers/facebook"
	"github.com/conku/auth/providers/github"
	"github.com/conku/auth/providers/google"
	"github.com/conku/auth/providers/twitter"
	"github.com/conku/auth_themes/clean"
	"github.com/conku/qor-example/config"
	"github.com/conku/qor-example/config/bindatafs"
	"github.com/conku/qor-example/config/db"
	"github.com/conku/qor-example/models/users"
	render "github.com/conku/render2"
)

var (
	// Auth initialize Auth for Authentication
	Auth = clean.New(&auth.Config{
		DB:         db.DB,
		Mailer:     config.Mailer,
		Render:     render.New(&render.Config{AssetFileSystem: bindatafs.AssetFS.NameSpace("auth")}),
		UserModel:  users.User{},
		Redirector: auth.Redirector{RedirectBack: config.RedirectBack},
	})

	// Authority initialize Authority for Authorization
	Authority = authority.New(&authority.Config{
		Auth: Auth,
	})
)

func init() {
	Auth.RegisterProvider(github.New(&config.Config.Github))
	Auth.RegisterProvider(google.New(&config.Config.Google))
	Auth.RegisterProvider(facebook.New(&config.Config.Facebook))
	Auth.RegisterProvider(twitter.New(&config.Config.Twitter))

	Authority.Register("logged_in_half_hour", authority.Rule{TimeoutSinceLastLogin: time.Minute * 30})
}
