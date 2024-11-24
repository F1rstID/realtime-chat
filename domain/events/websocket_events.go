// domain/events/websocket_events.go

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

// Common response codes
const (
	StatusSuccess = 2000
	StatusCreated = 2001
)

// WebSocketResponse represents the unified response structure for both REST and WebSocket
type WebSocketResponse struct {
	Success   bool        `json:"success"`
	Code      int         `json:"code"`
	Data      interface{} `json:"data"`
	Timestamp time.Time   `json:"timestamp"`
}

// MessageEventData represents the data structure for message events
type MessageEventData struct {
	Type           string    `json:"type"`
	MessageID      int       `json:"messageId"`
	ChatID         int       `json:"chatId"`
	SenderID       int       `json:"senderId"`
	SenderNickname string    `json:"senderNickname"`
	Content        string    `json:"content,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`
	UpdatedAt      time.Time `json:"updatedAt,omitempty"`
}

// NewWebSocketEvent creates a new WebSocket event with unified response format
func NewWebSocketEvent(eventType string, chatID int, data interface{}) *WebSocketResponse {
	eventData := MessageEventData{
		Type: eventType,
	}

	// Type assertion for different event types
	switch v := data.(type) {
	case *MessageEventData:
		eventData.MessageID = v.MessageID
		eventData.ChatID = v.ChatID
		eventData.SenderID = v.SenderID
		eventData.SenderNickname = v.SenderNickname
		eventData.Content = v.Content
		eventData.CreatedAt = v.CreatedAt
		eventData.UpdatedAt = v.UpdatedAt
	}

	return &WebSocketResponse{
		Success:   true,
		Code:      StatusSuccess,
		Data:      eventData,
		Timestamp: time.Now(),
	}
}

// ToJSON converts the WebSocket response to JSON bytes
func (r *WebSocketResponse) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}
