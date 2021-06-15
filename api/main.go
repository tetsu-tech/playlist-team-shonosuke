package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tetsu-tech/playlist-team-shonosuke/cmd/app"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
}

func main() {
	app.Run()
}
