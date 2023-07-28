package main

import "github.com/Vdolganov/shortify/internal/app/api"

func main() {
	server := api.GetNewServer()
	server.RunApp()
}
