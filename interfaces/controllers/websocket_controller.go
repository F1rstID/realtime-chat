// interfaces/controllers/websocket_controller.go
package controllers

import (
	"strconv"

	"github.com/f1rstid/realtime-chat/infrastructure/logger"
	"github.com/f1rstid/realtime-chat/infrastructure/websocket"
	"github.com/gofiber/fiber/v2"
	ws "github.com/gofiber/websocket/v2"
)

type WebSocketController struct {
	hub *websocket.Hub
}

func NewWebSocketController(hub *websocket.Hub) *WebSocketController {
	return &WebSocketController{
		hub: hub,
	}
}

// HandleWebSocket is a middleware for upgrading to websocket connections
func (wc *WebSocketController) HandleWebSocket(c *fiber.Ctx) error {
	if ws.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		c.Locals("userId", c.Locals("userId"))
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}

// WebSocket handles the WebSocket connection
func (wc *WebSocketController) WebSocket(c *ws.Conn) {
	// Get user ID from context (set by auth middleware)
	userID := c.Locals("userId")
	if userID == nil {
		c.Close()
		return
	}

	// Get chat ID from URL parameter
	chatIDStr := c.Params("chatId")
	if chatIDStr == "" {
		c.Close()
		return
	}

	chatID, err := strconv.Atoi(chatIDStr)
	if err != nil {
		logger.Error("Invalid chatId format: %v", err)
		c.Close()
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		logger.Error("Invalid userId type: %T", userID)
		c.Close()
		return
	}

	logger.Info("New WebSocket connection - UserID: %d, ChatID: %d", userIDInt, chatID)

	// Create new client
	client := &websocket.Client{
		Hub:    wc.hub,
		Conn:   c,
		Send:   make(chan []byte, 256),
		UserID: userIDInt,
		ChatID: chatID,
	}

	client.Hub.RegisterClient(client)

	// Setup ping handler to maintain connection
	c.SetCloseHandler(func(code int, text string) error {
		logger.Info("WebSocket connection closed - UserID: %d, ChatID: %d", userIDInt, chatID)
		return nil
	})

	// Start client message pumps
	go client.WritePump()
	client.ReadPump()
}
