package main

import (
	"github.com/Vdolganov/shortify/internal/app/app"
	"github.com/Vdolganov/shortify/internal/app/shorter"
	"github.com/Vdolganov/shortify/internal/app/storage/links"
	"github.com/Vdolganov/shortify/internal/config"
)

func main() {
	appConfig := config.New()
	st := links.New()
	sh := shorter.New(st)
	server := app.New(appConfig, sh)
	server.RunApp()
}
