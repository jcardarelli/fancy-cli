package main

import (
	"log"

	"github.com/jcardarelli/fancy-cli/cmd"
	db "github.com/jcardarelli/fancy-cli/database"
)

func main() {
	err := db.OpenConnectionPool("fancy-cli.db")
	if err != nil {
		log.Fatalln(err)
	}
	cmd.Execute()
}
