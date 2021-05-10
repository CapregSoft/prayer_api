package main

import "github.com/CapregSoft/prayer_api/server"

func main() {
	salah := server.New()
	salah.Start(":8080")

}
