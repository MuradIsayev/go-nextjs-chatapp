package main

import (
	"log"

	"github.com/MuradIsayev/go-nextjs-chatapp/db"
	"github.com/MuradIsayev/go-nextjs-chatapp/internal/user"
	"github.com/MuradIsayev/go-nextjs-chatapp/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could't initialize database connection: %s", err.Error())
	}

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	router.InitRouter(userHandler)
	router.Start("localhost:8080")

}
