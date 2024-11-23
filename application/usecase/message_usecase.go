package usecase

import (
	"errors"
	"time"

	"github.com/f1rstid/realtime-chat/domain/dto"
	"github.com/f1rstid/realtime-chat/domain/events"
	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
	"github.com/f1rstid/realtime-chat/infrastructure/logger"
	"github.com/f1rstid/realtime-chat/infrastructure/websocket"
)

type MessageUsecase struct {
	messageRepo repositories.MessageRepository
	chatRepo    repositories.ChatRepository
	wsHub       *websocket.Hub
}

func NewMessageUsecase(
	messageRepo repositories.MessageRepository,
	chatRepo repositories.ChatRepository,
	wsHub *websocket.Hub,
) *MessageUsecase {
	return &MessageUsecase{
		messageRepo: messageRepo,
		chatRepo:    chatRepo,
		wsHub:       wsHub,
	}
}

// SendMessage sends a new message in a chat
func (mu *MessageUsecase) SendMessage(chatID, senderID int, content string) (*dto.MessageResponse, error) {
	// Verify chat exists
	chat, err := mu.chatRepo.FindById(chatID)
	if err != nil {
		return nil, errors.New("chat not found")
	}

	message := &models.Message{
		ChatId:    chat.ID,
		SenderId:  senderID,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := mu.messageRepo.Create(message); err != nil {
		logger.Error("Failed to create message: %v", err)
		return nil, err
	}

	// Create and broadcast WebSocket event
	eventData := &events.MessageEventData{
		MessageID: message.ID,
		ChatID:    message.ChatId,
		SenderID:  message.SenderId,
		Content:   message.Content,
		CreatedAt: message.CreatedAt,
		UpdatedAt: message.UpdatedAt,
	}

	event := events.NewWebSocketEvent(events.EventMessageCreated, chatID, eventData)
	if eventJSON, err := event.ToJSON(); err == nil {
		mu.wsHub.BroadcastToChat(chatID, eventJSON)
	}

	return dto.NewMessageResponse(message), nil
}

// UpdateMessage updates an existing message
func (mu *MessageUsecase) UpdateMessage(messageID, userID int, newContent string) (*dto.MessageResponse, error) {
	// 기존 메시지를 먼저 조회
	originalMessage, err := mu.messageRepo.FindById(messageID)
	if err != nil {
		return nil, errors.New("message not found")
	}

	if originalMessage.SenderId != userID {
		return nil, errors.New("unauthorized to update this message")
	}

	// 메시지 업데이트
	updatedMessage := &models.Message{
		ID:        originalMessage.ID,
		ChatId:    originalMessage.ChatId,
		SenderId:  originalMessage.SenderId,
		Content:   newContent,
		CreatedAt: originalMessage.CreatedAt, // 원본 생성 시간 유지
		UpdatedAt: time.Now(),
	}

	if err := mu.messageRepo.Update(updatedMessage); err != nil {
		return nil, err
	}

	// Create and broadcast WebSocket event
	eventData := &events.MessageEventData{
		MessageID: updatedMessage.ID,
		ChatID:    updatedMessage.ChatId,
		SenderID:  updatedMessage.SenderId,
		Content:   updatedMessage.Content,
		CreatedAt: updatedMessage.CreatedAt, // 원본 생성 시간 포함
		UpdatedAt: updatedMessage.UpdatedAt,
	}

	event := events.NewWebSocketEvent(events.EventMessageUpdated, updatedMessage.ChatId, eventData)
	if eventJSON, err := event.ToJSON(); err == nil {
		mu.wsHub.BroadcastToChat(updatedMessage.ChatId, eventJSON)
	}

	return dto.NewMessageResponse(updatedMessage), nil
}

// DeleteMessage deletes an existing message
func (mu *MessageUsecase) DeleteMessage(messageID, userID int) error {
	message, err := mu.messageRepo.FindById(messageID)
	if err != nil {
		return errors.New("message not found")
	}

	if message.SenderId != userID {
		return errors.New("unauthorized to delete this message")
	}

	if err := mu.messageRepo.Delete(messageID); err != nil {
		return err
	}

	// Create and broadcast WebSocket event
	eventData := &events.MessageEventData{
		MessageID: message.ID,
		ChatID:    message.ChatId,
		SenderID:  message.SenderId,
		CreatedAt: message.CreatedAt,
		UpdatedAt: time.Now(),
	}

	event := events.NewWebSocketEvent(events.EventMessageDeleted, message.ChatId, eventData)
	if eventJSON, err := event.ToJSON(); err == nil {
		mu.wsHub.BroadcastToChat(message.ChatId, eventJSON)
	}

	return nil
}
