package main

import (
	"log"
	// "net/http"
	// "sync"

	"github.com/gin-gonic/gin"
	// "github.com/gorilla/websocket"

	"chat-app/controllers"
)



func main() {
	r := gin.Default()

	// Serve the homepage
	r.StaticFile("/", "./static/index.html")

	// WebSocket route
	r.GET("/ws", controllers.WebSocketHandler)

	// Start the server
	log.Println("Server started on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error starting server:", err)
	}
}




