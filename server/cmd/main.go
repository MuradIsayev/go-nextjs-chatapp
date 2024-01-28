package main

import (
	"log"

	"github.com/MuradIsayev/go-nextjs-chatapp/db"
)

func main() {
	_, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could't initialize database connection: %s", err.Error())
	}
}
