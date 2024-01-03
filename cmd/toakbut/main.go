package main

import (
	"fmt"
	"toakbut/cmd/toakbut/config"
	"toakbut/cmd/toakbut/server"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	cfg := config.Config{}
	envconfig.MustProcess("API", &cfg)
	fmt.Println("starting server : ", cfg.Name)
	app := toakbut.InitServer(&cfg)
	startServerString := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	err := app.Listen(startServerString)
	if err != nil {
		panic(err)
	}
}
