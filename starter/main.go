package main

import "github.com/CapregSoft/prayer-api/server"

func main() {
	salah := server.New()
	salah.Start(":8080")

}
