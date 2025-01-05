package routers

import (
	"github.com/f1rstid/realtime-chat/application/usecase"
	"github.com/f1rstid/realtime-chat/config"
	"github.com/f1rstid/realtime-chat/domain/services"
	"github.com/f1rstid/realtime-chat/infrastructure/sqlite"
	"github.com/f1rstid/realtime-chat/infrastructure/websocket"
	"github.com/f1rstid/realtime-chat/interfaces/controllers"
	"github.com/f1rstid/realtime-chat/interfaces/middlewares"
	"github.com/f1rstid/realtime-chat/interfaces/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	ws "github.com/gofiber/websocket/v2"
)

func SetRoutes(app *fiber.App, config *config.Config) {
	// Initialize WebSocket hub
	wsHub := websocket.NewHub()
	go wsHub.Run()

	// Initialize repositories
	userRepo := repositories.NewUserRepository(sqlite.DB)
	chatRepo := repositories.NewChatRepository(sqlite.DB)
	messageRepo := repositories.NewMessageRepository(sqlite.DB)

	// Initialize services
	authService := services.NewAuthService(config.JWTSecret)
	userService := services.NewUserService(userRepo)

	// Initialize usecases
	authUseCase := usecase.NewAuthUsecase(userRepo, authService)
	chatUseCase := usecase.NewChatUsecase(chatRepo, messageRepo, userRepo)
	messageUseCase := usecase.NewMessageUsecase(messageRepo, chatRepo, wsHub)
	userUseCase := usecase.NewUserUseCase(userRepo, userService)

	// Initialize controllers
	authController := controllers.NewAuthController(authUseCase)
	chatController := controllers.NewChatController(chatUseCase, messageUseCase)
	messageController := controllers.NewMessageController(messageUseCase)
	wsController := controllers.NewWebSocketController(wsHub)
	userController := controllers.NewUserController(userUseCase)

	// Swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Auth routes
	auth := app.Group("/api/auth")
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	// Protected routes
	api := app.Group("/api", middlewares.AuthMiddleware(authService))

	// Chat routes
	chats := api.Group("/chats")
	chats.Get("/", chatController.GetUserChats)
	chats.Post("/private", chatController.CreatePrivateChat)
	chats.Post("/group", chatController.CreateGroupChat)
	api.Get("/chats/:chatId/messages", messageController.GetChatMessages)

	// Message routes
	messages := api.Group("/messages")
	messages.Post("/", messageController.SendMessage)
	messages.Put("/:id", messageController.UpdateMessage)
	messages.Delete("/:id", messageController.DeleteMessage)

	users := api.Group("/users")
	users.Get("/", userController.GetAllUsers) // 새로운 라우트 추가

	// WebSocket routes with authentication
	//app.Use("/ws", middlewares.WebSocketAuthMiddleware(authService))
	//app.Use("/ws/:chatId", wsController.HandleWebSocket)
	//app.Get("/ws/:chatId", ws.New(wsController.WebSocket))
	app.Use("/ws", middlewares.WebSocketAuthMiddleware(authService))
	app.Get("/ws", ws.New(wsController.WebSocket))
}
