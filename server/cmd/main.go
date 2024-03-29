package main

import (
	"log"
	"os"

	"github.com/MuradIsayev/go-nextjs-chatapp/db"
	"github.com/MuradIsayev/go-nextjs-chatapp/internal/user"
	"github.com/MuradIsayev/go-nextjs-chatapp/internal/ws"
	"github.com/MuradIsayev/go-nextjs-chatapp/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbURL := os.Getenv("POSTGRES_URL")

	dbConn, err := db.NewDatabase(dbURL)
	if err != nil {
		log.Fatalf("Could't initialize database connection: %s", err.Error())
	}

	userRepository := user.NewRepository(dbConn.GetDB())
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	router.Start("localhost:8080")
}
