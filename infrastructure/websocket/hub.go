// infrastructure/websocket/hub.go
package websocket

import (
	"github.com/f1rstid/realtime-chat/infrastructure/logger"
	"github.com/gofiber/websocket/v2"
	"sync"
)

// Hub maintains the set of active clients and broadcasts messages
type Hub struct {
	// Registered clients mapped by user ID
	clients map[int]map[*Client]bool

	// Register requests from the clients
	register chan *Client

	// Unregister requests from clients
	unregister chan *Client

	// Mutex for thread-safe operations on the clients map
	mu sync.RWMutex
}

// Client represents a connected websocket client
type Client struct {
	Hub    *Hub
	Conn   *websocket.Conn
	Send   chan []byte
	UserID int
}

// NewHub creates a new Hub instance
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[int]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Run starts the hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if _, ok := h.clients[client.UserID]; !ok {
				h.clients[client.UserID] = make(map[*Client]bool)
			}
			h.clients[client.UserID][client] = true
			h.mu.Unlock()
			logger.Info("Client registered - UserID: %d", client.UserID)

		case client := <-h.unregister:
			h.mu.Lock()
			if clients, ok := h.clients[client.UserID]; ok {
				if _, ok := clients[client]; ok {
					delete(clients, client)
					close(client.Send)
					if len(clients) == 0 {
						delete(h.clients, client.UserID)
					}
				}
			}
			h.mu.Unlock()
			logger.Info("Client unregistered - UserID: %d", client.UserID)
		}
	}
}

// BroadcastToUsers sends a message to specified users
func (h *Hub) BroadcastToUsers(userIDs []int, message []byte) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	for _, userID := range userIDs {
		if clients, ok := h.clients[userID]; ok {
			for client := range clients {
				select {
				case client.Send <- message:
					logger.Info("Message sent to UserID: %d", userID)
				default:
					close(client.Send)
					delete(clients, client)
					if len(clients) == 0 {
						delete(h.clients, userID)
					}
					logger.Error("Failed to send message to UserID: %d", userID)
				}
			}
		}
	}
}

// RegisterClient adds a new client to the hub
func (h *Hub) RegisterClient(client *Client) {
	h.register <- client
}

// UnregisterClient removes a client from the hub
func (h *Hub) UnregisterClient(client *Client) {
	h.unregister <- client
}

// WritePump pumps messages from the hub to the websocket connection
func (c *Client) WritePump() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.Conn.WriteMessage(websocket.TextMessage, message); err != nil {
				logger.Error("Failed to send message to client %d: %v", c.UserID, err)
				return
			}
		}
	}
}

// ReadPump pumps messages from the websocket connection to the hub
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.UnregisterClient(c)
		c.Conn.Close()
	}()

	for {
		_, _, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logger.Error("WebSocket read error: %v", err)
			}
			break
		}
	}
}
