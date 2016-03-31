package config

import (
	"github.com/jinzhu/configor"
	"github.com/qor/render"
)

var Render *render.Render
var Config = struct {
	Port uint `default:"7000" env:"PORT"`
	DB   struct {
		Name     string
		Adapter  string `default:"sqlite"`
		User     string
		Password string
		Host     string `default:"localhost"`
	}
}{}

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/application.yml"); err != nil {
		panic(err)
	}

	Render = render.New()
}
