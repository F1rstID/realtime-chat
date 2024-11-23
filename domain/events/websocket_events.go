package events

import (
	"encoding/json"
	"time"
)

// Event types
const (
	EventMessageCreated = "message.created"
	EventMessageUpdated = "message.updated"
	EventMessageDeleted = "message.deleted"
)

// WebSocketEvent represents the structure of a WebSocket event
type WebSocketEvent struct {
	Type      string      `json:"type"`
	ChatID    int         `json:"chatId"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

// MessageEventData represents the data structure for message events
type MessageEventData struct {
	MessageID int       `json:"messageId"`
	ChatID    int       `json:"chatId"`
	SenderID  int       `json:"senderId"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// NewWebSocketEvent creates a new WebSocket event
func NewWebSocketEvent(eventType string, chatID int, data interface{}) *WebSocketEvent {
	return &WebSocketEvent{
		Type:      eventType,
		ChatID:    chatID,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// ToJSON converts the WebSocket event to JSON bytes
func (e *WebSocketEvent) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}
