package main

import (
	"ChatRoooms/server/db"
	"ChatRoooms/server/internal/router"
	"ChatRoooms/server/internal/user"
	"ChatRoooms/server/internal/ws"
	"log"
)

func main() {
	Database, err := db.NewDatabase()
	if err != nil {
		log.Fatalln("db connection not successful")
		return
	}

	UserRepository := user.NewRepository(Database.GetDB())
	UserService := user.NewService(&UserRepository)
	UserHandler := user.NewHandler(&UserService)

	Hub := ws.NewHub()
	WsHandler := ws.NewHandler(Hub)
	go Hub.Run()

	router.InitRoutes(UserHandler, *WsHandler)
	router.Start(":8080")

}
