package main

import (
	"github.com/Vdolganov/shortify/internal/app/api"
	"github.com/Vdolganov/shortify/internal/config"
)

func main() {
	appConfig := config.InitConfig()
	server := api.GetNewServer(appConfig.ServerAddress, appConfig.BaseURL)
	server.RunApp()
}
