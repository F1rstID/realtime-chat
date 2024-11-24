package main

import (
	"github.com/f1rstid/realtime-chat/docs"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/f1rstid/realtime-chat/config"
	_ "github.com/f1rstid/realtime-chat/docs"
	"github.com/f1rstid/realtime-chat/infrastructure/logger"
	"github.com/f1rstid/realtime-chat/infrastructure/sqlite"
	"github.com/f1rstid/realtime-chat/interfaces/middlewares"
	"github.com/f1rstid/realtime-chat/interfaces/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// @title           Realtime Chat API
// @version         1.0
// @description     실시간 채팅을 위한 RESTful API 서버입니다.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:5050
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description 'Bearer ' 접두사와 함께 JWT 토큰을 입력하세요. 예시: "Bearer eyJhbGciOi..."

// @Security Bearer
func main() {
	// 설정 로드
	config, err := config.LoadConfig()
	if err != nil {
		logger.Error("Failed to load config: %v", err)
		log.Fatal(err)
	}

	// 데이터베이스 디렉토리 확인 및 생성
	dbDir := filepath.Dir(config.Database.DSN)
	if err := os.MkdirAll(dbDir, 0755); err != nil {
		logger.Error("Failed to create database directory: %v", err)
		log.Fatal(err)
	}

	// 데이터베이스 초기화
	if err := sqlite.InitDB(config.Database.DSN); err != nil {
		logger.Error("Failed to initialize database: %v", err)
		log.Fatal(err)
	}

	// 데이터베이스 마이그레이션
	if err := sqlite.Migrate(); err != nil {
		logger.Error("Failed to migrate database: %v", err)
		log.Fatal(err)
	}
	defer sqlite.CloseDB()

	// Fiber 앱 생성
	app := fiber.New(fiber.Config{
		ErrorHandler: middlewares.ErrorHandler(),
	})

	// 미들웨어 설정
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:  "*",
		AllowHeaders:  "*",
		AllowMethods:  "*",
		ExposeHeaders: "*",
	}))
	app.Use(middlewares.RequestLogger())

	docs.SwaggerInfo.Host = config.ServerURL + ":" + config.ServerPort

	// Swagger 설정
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:         "/swagger/doc.json",
		DeepLinking: true,
	}))

	// 라우터 설정
	routers.SetRoutes(app, config)

	// 서버 시작
	go func() {
		logger.Info("Server is starting on port %s", config.ServerPort)
		logger.Info("Swagger documentation is available at http://localhost:%s/swagger/", config.ServerPort)
		if err := app.Listen(":" + config.ServerPort); err != nil {
			logger.Error("Server failed to start: %v", err)
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		logger.Error("Failed to shutdown server: %v", err)
		log.Fatal(err)
	}
}
