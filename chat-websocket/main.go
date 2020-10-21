package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	hub := newServeWs()
	go hub.run()
	roomHub := newRoom()
	go roomHub.run()

	router := gin.Default()
	router.LoadHTMLFiles("index.html")

	// add new user to socket
	router.GET("/users/:userId", func(c *gin.Context) {
		userID := c.Param("userId")
		hub.serveWsForAddUser(c.Writer, c.Request, userID)
		// c.HTML(200, "index.html", nil)
	})

	router.GET("/users/:userId/chat/:roomId", func(c *gin.Context) {
		userID := c.Param("userId")
		roomID := c.Param("roomId")
		hub.serveWsForAssignUserToRoom(c.Writer, c.Request, userID, roomID)
	})

	// assign user to room
	router.POST("/room/:roomId/user/:userId", func(c *gin.Context) {
		userID := c.Param("userId")
		roomID := c.Param("roomId")
		hub.serveWsForAssignUserToRoom(c.Writer, c.Request, userID, roomID)
		// c.HTML(200, "index.html", nil)
	})

	// add new room to socket
	router.POST("/room/:roomId", func(c *gin.Context) {
		roomID := c.Param("roomId")
		hub.serveWsForAddRoom(c.Writer, c.Request, roomID)
		// c.HTML(200, "index.html", nil)
	})

	router.Run(fmt.Sprintf(":%d", 8080))
}
