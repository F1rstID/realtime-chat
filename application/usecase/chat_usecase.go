package usecase

import (
	"errors"
	"github.com/f1rstid/realtime-chat/domain/dto"
	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
	"log"
)

type ChatUsecase struct {
	chatRepo    repositories.ChatRepository
	messageRepo repositories.MessageRepository
	userRepo    repositories.UserRepository
}

func NewChatUsecase(
	chatRepo repositories.ChatRepository,
	messageRepo repositories.MessageRepository,
	userRepo repositories.UserRepository,
) *ChatUsecase {
	return &ChatUsecase{
		chatRepo:    chatRepo,
		messageRepo: messageRepo,
		userRepo:    userRepo,
	}
}

// GetUserChats returns all chats for a user with their last messages
func (cu *ChatUsecase) GetUserChats(userID int) ([]dto.ChatResponse, error) {
	// Verify user exists
	_, err := cu.userRepo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Get user's chats
	chats, err := cu.chatRepo.GetUserChats(userID)
	if err != nil {
		return nil, err
	}

	// Get chat IDs
	chatIDs := make([]int, len(chats))
	for i, chat := range chats {
		chatIDs[i] = chat.ID
	}

	// Get last messages for all chats
	lastMessages, err := cu.chatRepo.GetLastMessages(chatIDs)
	if err != nil {
		log.Println("Failed to get last messages:", err)
		return nil, err
	}

	// Create response
	return dto.NewChatListResponse(chats, lastMessages), nil
}

func (cu *ChatUsecase) CreatePrivateChat(user1ID, user2ID int) (*models.Chat, error) {
	// Verify both users exist
	user1, err := cu.userRepo.FindByID(user1ID)
	if err != nil {
		return nil, errors.New("user1 not found")
	}

	user2, err := cu.userRepo.FindByID(user2ID)
	if err != nil {
		return nil, errors.New("user2 not found")
	}

	chat := &models.Chat{
		Name: user1.Nickname + "-" + user2.Nickname,
	}

	if err := cu.chatRepo.Create(chat); err != nil {
		return nil, err
	}

	if err := cu.chatRepo.AddUserToChat(chat.ID, user1ID); err != nil {
		cu.chatRepo.Delete(chat.ID)
		return nil, errors.New("failed to add user1 to chat")
	}

	if err := cu.chatRepo.AddUserToChat(chat.ID, user2ID); err != nil {
		cu.chatRepo.Delete(chat.ID)
		return nil, errors.New("failed to add user2 to chat")
	}
	return chat, nil
}

func (cu *ChatUsecase) CreateGroupChat(name string, userIDs []int) (*models.Chat, error) {
	// Verify all users exist
	for _, userID := range userIDs {
		_, err := cu.userRepo.FindByID(userID)
		if err != nil {
			return nil, errors.New("user not found: " + string(userID))
		}
	}

	chat := &models.Chat{
		Name: name,
	}

	if err := cu.chatRepo.Create(chat); err != nil {
		return nil, err
	}

	// Add users to the chat group
	for _, userID := range userIDs {
		if err := cu.chatRepo.AddUserToChat(chat.ID, userID); err != nil {
			cu.chatRepo.Delete(chat.ID)
			return nil, errors.New("failed to add user to chat group")
		}
	}

	return chat, nil
}
