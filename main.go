package main

import (
	"log"
	// "net/http"
	// "sync"

	"github.com/gin-gonic/gin"
	// "github.com/gorilla/websocket"

	"chat-app/controllers"
)

// Upgrader to upgrade HTTP requests to WebSocket connections
// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true
// 	},
// }

// Structures to manage clients and groups
// var clients = make(map[string]*websocket.Conn) // Map user ID to WebSocket connection
// var groups = make(map[string]map[string]*websocket.Conn) // Map group ID to users in the group
// var mu sync.Mutex // Mutex to manage concurrent access

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

// Handle WebSocket connections
// func handleConnections(c *gin.Context) {
// 	userID := c.Query("userID")    // Get user ID from query parameters
// 	groupID := c.Query("groupID")  // Get group ID from query parameters

// 	if userID == "" || groupID == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "userID and groupID are required"})
// 		return
// 	}

// 	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
// 	if err != nil {
// 		log.Println("WebSocket Upgrade Error:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Add connection to clients and groups
// 	mu.Lock()
// 	clients[userID] = conn
// 	if groups[groupID] == nil {
// 		groups[groupID] = make(map[string]*websocket.Conn)
// 	}
// 	groups[groupID][userID] = conn
// 	mu.Unlock()

// 	log.Printf("User %s joined group %s", userID, groupID)

// 	for {
// 		// Read message from client
// 		messageType, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Printf("Error reading message from %s: %v", userID, err)
// 			removeClient(userID, groupID)
// 			break
// 		}

// 		// Broadcast message to the group
// 		broadcastToGroup(groupID, userID, messageType, msg)
// 	}
// }

// // Remove a user from clients and groups
// func removeClient(userID, groupID string) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	if _, exists := clients[userID]; exists {
// 		delete(clients, userID)
// 	}
// 	if _, exists := groups[groupID]; exists {
// 		delete(groups[groupID], userID)
// 		if len(groups[groupID]) == 0 {
// 			delete(groups, groupID)
// 		}
// 	}
// 	log.Printf("User %s left group %s", userID, groupID)
// }

// // Broadcast a message to all users in a group except the sender
// func broadcastToGroup(groupID, senderID string, messageType int, message []byte) {
// 	mu.Lock()
// 	defer mu.Unlock()

// 	for userID, conn := range groups[groupID] {
// 		// if userID != senderID { // Avoid sending the message to the sender
// 			if err := conn.WriteMessage(messageType, message); err != nil {
// 				log.Printf("Error sending message to %s: %v", userID, err)
// 				conn.Close()
// 				delete(groups[groupID], userID)
// 			}
// 		// }
// 	}
// }
