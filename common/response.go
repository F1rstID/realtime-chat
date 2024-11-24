package common

// Response codes
const (
	// Success codes (2xxx)
	StatusSuccess = 2000
	StatusCreated = 2001

	// Client errors (4xxx)
	StatusInvalidRequest  = 4000
	StatusUnauthorized    = 4001
	StatusForbidden       = 4002
	StatusNotFound        = 4003
	StatusEmailExists     = 4004
	StatusNicknameExists  = 4005
	StatusInvalidAuth     = 4006
	StatusChatNotFound    = 4007
	StatusMessageNotFound = 4008
	StatusUnauthorizedMsg = 4009

	// Server errors (5xxx)
	StatusInternalError = 5000
	StatusDatabaseError = 5001
)

// BaseResponse 기본 응답 구조
type BaseResponse struct {
	Success bool        `json:"success" example:"true"`
	Code    int         `json:"code" example:"2000"`
	Data    interface{} `json:"data"`
}

// ErrorResponse 에러 응답 구조체
type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4000"`
	Data    string `json:"data" example:"잘못된 요청입니다"`
}

// UserData represents user information
type UserData struct {
	ID        int    `json:"id" example:"1"`
	Email     string `json:"email" example:"user@example.com"`
	Nickname  string `json:"nickname" example:"홍길동"`
	CreatedAt string `json:"createdAt" example:"2024-03-23T12:00:00Z"`
}

// AuthData represents authentication data
type AuthData struct {
	Token string   `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
	User  UserData `json:"user"`
}

// UserInfo in chat room
type UserInfo struct {
	ID       int    `json:"id" example:"1"`
	Nickname string `json:"nickname" example:"홍길동"`
}

// LastMessage represents last message in chat
type LastMessage struct {
	Content        string `json:"content" example:"안녕하세요"`
	SenderID       int    `json:"senderId" example:"1"`
	SenderNickname string `json:"senderNickname" example:"홍길동"`
	CreatedAt      string `json:"createdAt" example:"2024-03-23T12:00:00Z"`
}

// ChatData represents basic chat information (for creation/update)
type ChatData struct {
	ID        int    `json:"id" example:"1"`
	Name      string `json:"name" example:"개발팀 채팅방"`
	CreatedAt string `json:"createdAt" example:"2024-03-23T12:00:00Z"`
}

// ChatListData represents chat information with users (for list view)
type ChatListData struct {
	ID          int          `json:"id" example:"1"`
	Name        string       `json:"name" example:"개발팀 채팅방"`
	CreatedAt   string       `json:"createdAt" example:"2024-03-23T12:00:00Z"`
	LastMessage *LastMessage `json:"lastMessage,omitempty"`
	Users       []UserInfo   `json:"users"`
}

// MessageData represents message information
type MessageData struct {
	ID             int    `json:"id" example:"1"`
	ChatID         int    `json:"chatId" example:"1"`
	SenderID       int    `json:"senderId" example:"1"`
	SenderNickname string `json:"senderNickname" example:"홍길동"`
	Content        string `json:"content" example:"안녕하세요"`
	CreatedAt      string `json:"createdAt" example:"2024-03-23T12:00:00Z"`
	UpdatedAt      string `json:"updatedAt" example:"2024-03-23T12:00:00Z"`
}

// Predefined errors
var (
	// ErrInvalidRequest 잘못된 요청
	InvalidRequest = ErrorResponse{
		Success: false,
		Code:    StatusInvalidRequest,
		Data:    "잘못된 요청입니다",
	}

	// ErrUnauthorized 인증 필요
	Unauthorized = ErrorResponse{
		Success: false,
		Code:    StatusUnauthorized,
		Data:    "인증이 필요합니다",
	}

	// ErrInvalidAuth 인증 실패
	InvalidAuth = ErrorResponse{
		Success: false,
		Code:    StatusInvalidAuth,
		Data:    "이메일 또는 비밀번호가 올바르지 않습니다",
	}

	// ErrEmailExists 이미 존재하는 이메일
	EmailExists = ErrorResponse{
		Success: false,
		Code:    StatusEmailExists,
		Data:    "이미 사용중인 이메일입니다",
	}

	// ErrNicknameExists 이미 존재하는 닉네임
	NicknameExists = ErrorResponse{
		Success: false,
		Code:    StatusNicknameExists,
		Data:    "이미 사용중인 닉네임입니다",
	}

	// ErrChatNotFound 채팅방을 찾을 수 없음
	ChatNotFound = ErrorResponse{
		Success: false,
		Code:    StatusChatNotFound,
		Data:    "채팅방을 찾을 수 없습니다",
	}

	// ErrMessageNotFound 메시지를 찾을 수 없음
	MessageNotFound = ErrorResponse{
		Success: false,
		Code:    StatusMessageNotFound,
		Data:    "메시지를 찾을 수 없습니다",
	}

	// ErrUnauthorizedMessage 메시지에 대한 권한 없음
	UnauthorizedMessage = ErrorResponse{
		Success: false,
		Code:    StatusUnauthorizedMsg,
		Data:    "메시지에 대한 권한이 없습니다",
	}

	// ErrInternalServer 내부 서버 오류
	InternalServer = ErrorResponse{
		Success: false,
		Code:    StatusInternalError,
		Data:    "내부 서버 오류가 발생했습니다",
	}

	// ErrDatabase 데이터베이스 오류
	DatabaseError = ErrorResponse{
		Success: false,
		Code:    StatusDatabaseError,
		Data:    "데이터베이스 오류가 발생했습니다",
	}
)

// Swagger examples
type ErrInvalidRequest struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4000"`
	Data    string `json:"data" example:"잘못된 요청입니다"`
}

type ErrUnauthorized struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4001"`
	Data    string `json:"data" example:"인증이 필요합니다"`
}

type ErrInvalidAuth struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4006"`
	Data    string `json:"data" example:"이메일 또는 비밀번호가 올바르지 않습니다"`
}

type ErrEmailExists struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4004"`
	Data    string `json:"data" example:"이미 사용중인 이메일입니다"`
}

type ErrNicknameExists struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4005"`
	Data    string `json:"data" example:"이미 사용중인 닉네임입니다"`
}

type ErrChatNotFound struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4007"`
	Data    string `json:"data" example:"채팅방을 찾을 수 없습니다"`
}

type ErrMessageNotFound struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4008"`
	Data    string `json:"data" example:"메시지를 찾을 수 없습니다"`
}

type ErrUnauthorizedMessage struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"4009"`
	Data    string `json:"data" example:"메시지에 대한 권한이 없습니다"`
}

type ErrInternalServer struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"5000"`
	Data    string `json:"data" example:"내부 서버 오류가 발생했습니다"`
}

type ErrDatabase struct {
	Success bool   `json:"success" example:"false"`
	Code    int    `json:"code" example:"5001"`
	Data    string `json:"data" example:"데이터베이스 오류가 발생했습니다"`
}

// Success response examples
type RegisterResponse struct {
	Success bool     `json:"success" example:"true"`
	Code    int      `json:"code" example:"2001"`
	Data    AuthData `json:"data"`
}

type LoginResponse struct {
	Success bool     `json:"success" example:"true"`
	Code    int      `json:"code" example:"2000"`
	Data    AuthData `json:"data"`
}

type ChatResponse struct {
	Success bool     `json:"success" example:"true"`
	Code    int      `json:"code" example:"2000"`
	Data    ChatData `json:"data"`
}

type ChatListResponse struct {
	Success bool           `json:"success" example:"true"`
	Code    int            `json:"code" example:"2000"`
	Data    []ChatListData `json:"data"`
}

type MessageResponse struct {
	Success bool        `json:"success" example:"true"`
	Code    int         `json:"code" example:"2000"`
	Data    MessageData `json:"data"`
}

type MessageListData struct {
	ChatId        int           `json:"chatId" example:"1"`
	Messages      []MessageData `json:"messages"`
	LastMessageId int           `json:"lastMessageId" example:"100"`
	HasMore       bool          `json:"hasMore" example:"true"`
	NextCursor    int           `json:"nextCursor" example:"50"`
}

type MessageListResponse struct {
	Success bool            `json:"success" example:"true"`
	Code    int             `json:"code" example:"2000"`
	Data    MessageListData `json:"data"`
}

type CreateChatRequest struct {
	Name    string `json:"name" example:"Team Chat" validate:"required"`
	UserIDs []int  `json:"user_ids" example:"[1,2,3]" validate:"required"`
}

type CreatePrivateChatRequest struct {
	TargetId int `json:"targetId" example:"1"`
}
