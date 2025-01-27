package main

import (
	"log"
	"tidy/db"
	"tidy/internal/user"
	"tidy/internal/websocket"
	"tidy/router"
)

func main() {
	dbConn, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialise db connection %s", err)
	}

	repo := user.NewRepository(dbConn.GetDB())
	service := user.NewService(&repo)
	userHandler := user.NewHandler(&service)

	hub := websocket.NewHub()
	WsHandler := websocket.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, WsHandler)
	router.Start("localhost:3000")
}
