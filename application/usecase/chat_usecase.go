package usecase

import (
	"errors"
	"fmt"
	"github.com/f1rstid/realtime-chat/domain/dto"
	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
	"github.com/f1rstid/realtime-chat/infrastructure/logger"
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

// GetUserChats returns all chats for a user with their last messages and users
func (cu *ChatUsecase) GetUserChats(userID int) ([]dto.ChatListResponse, error) {
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
		logger.Error("Failed to get last messages: %v", err)
		return nil, err
	}

	// Get users for all chats
	usersMap := make(map[int][]models.User)
	for _, chatID := range chatIDs {
		users, err := cu.chatRepo.GetChatUsers(chatID)
		if err != nil {
			logger.Error("Failed to get chat users for chatID %d: %v", chatID, err)
			continue
		}
		usersMap[chatID] = users
	}

	// Create response
	return dto.NewChatListResponse(chats, lastMessages, usersMap), nil
}

func (cu *ChatUsecase) CreatePrivateChat(user1ID, user2ID int) (*dto.ChatResponse, error) {
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

	return dto.NewChatResponse(chat), nil
}

func (cu *ChatUsecase) CreateGroupChat(name string, userIDs []int) (*dto.ChatResponse, error) {
	// Check for empty name
	if name == "" {
		return nil, errors.New("chat name is required")
	}

	// Make sure we have at least one other user
	if len(userIDs) < 1 {
		return nil, errors.New("at least one other user is required")
	}

	// Create a map for O(1) lookup to check for duplicates
	userIDMap := make(map[int]bool)
	uniqueUserIDs := []int{}

	// Add all users to the map, excluding duplicates
	for _, userID := range userIDs {
		if !userIDMap[userID] {
			userIDMap[userID] = true
			uniqueUserIDs = append(uniqueUserIDs, userID)
		}
	}

	// Verify all users exist
	for _, userID := range uniqueUserIDs {
		_, err := cu.userRepo.FindByID(userID)
		if err != nil {
			return nil, fmt.Errorf("user not found: %d", userID)
		}
	}

	chat := &models.Chat{
		Name: name,
	}

	if err := cu.chatRepo.Create(chat); err != nil {
		return nil, err
	}

	// Add all unique users to the chat group
	for _, userID := range uniqueUserIDs {
		if err := cu.chatRepo.AddUserToChat(chat.ID, userID); err != nil {
			cu.chatRepo.Delete(chat.ID)
			return nil, errors.New("failed to add user to chat group")
		}
	}

	return dto.NewChatResponse(chat), nil
}
