package router

import (
	"ChatRoooms/server/internal/user"
	"ChatRoooms/server/internal/ws"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRoutes(userHandler user.Handler, wsHandler ws.Handler) *gin.Engine {
	router = gin.Default()

	router.POST("/sign-up", userHandler.CreateUser)
	router.GET("/sign-in", userHandler.Login)
	router.GET("/logout", userHandler.Logout)

	router.POST("/ws/createRoom", wsHandler.CreateRoom)
	router.GET("/ws/joinRoom/:roomID", wsHandler.JoinRoom)
	router.GET("/ws/getRooms", wsHandler.GetRooms)
	router.GET("/ws/getClients/:roomID", wsHandler.GetClients)
	return router
}

func Start(addr string) error {
	return router.Run(addr)
}
