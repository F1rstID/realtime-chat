package usecase

import (
	"errors"
	"github.com/f1rstid/realtime-chat/domain/models"
	"github.com/f1rstid/realtime-chat/domain/repositories"
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

// CreatePrivateChat creates a 1:1 chat between two users
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
			// If there's an error, we might want to clean up the created chat
			cu.chatRepo.Delete(chat.ID)
			return nil, errors.New("failed to add user to chat group")
		}
	}

	return chat, nil
}

// GetUserChats returns all chats for a user
func (cu *ChatUsecase) GetUserChats(userID int) ([]models.Chat, error) {
	// Implementation here
	return nil, nil
}
