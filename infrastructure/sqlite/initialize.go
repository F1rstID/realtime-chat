package sqlite

import (
	"github.com/gofiber/websocket/v2"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sqlx.DB

func InitDB(dataSource string) error {
	var err error
	DB, err = sqlx.Open("sqlite3", dataSource)
	if err != nil {
		return err
	}

	// Set connection pool settings
	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(25)
	DB.SetConnMaxLifetime(5 * time.Minute)

	err = DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func Migrate() error {
	sql := `
	-- Users table
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL,
		nickname TEXT NOT NULL UNIQUE,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- Chats table
	CREATE TABLE IF NOT EXISTS chats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	-- Messages table
	CREATE TABLE IF NOT EXISTS messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		chatId INTEGER NOT NULL,
		senderId INTEGER NOT NULL,
		content TEXT NOT NULL,
		createdAt DATETIME NOT NULL,
		updatedAt DATETIME NOT NULL,
		FOREIGN KEY (chatId) REFERENCES chats(id) ON DELETE CASCADE,
		FOREIGN KEY (senderId) REFERENCES users(id) ON DELETE CASCADE
	);

	-- Chat groups table (for group chats)
	CREATE TABLE IF NOT EXISTS chat_groups (
		chatId INTEGER NOT NULL,
		userId INTEGER NOT NULL,
		createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (chatId, userId),
		FOREIGN KEY (chatId) REFERENCES chats(id) ON DELETE CASCADE,
		FOREIGN KEY (userId) REFERENCES users(id) ON DELETE CASCADE
	);

	-- Create indexes
	CREATE INDEX IF NOT EXISTS idx_messages_chatId ON messages(chatId);
	CREATE INDEX IF NOT EXISTS idx_messages_senderId ON messages(senderId);
	CREATE INDEX IF NOT EXISTS idx_chat_groups_chatId ON chat_groups(chatId);
	CREATE INDEX IF NOT EXISTS idx_chat_groups_userId ON chat_groups(userId);
	`

	_, err := DB.Exec(sql)
	if err != nil {
		return err
	}

	log.Println("Database migration completed successfully")
	return nil
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}

// WebSocket hub for managing connections and broadcasting messages
type Hub struct {
	// Registered clients
	clients map[*Client]bool

	// Register requests from clients
	register chan *Client

	// Unregister requests from clients
	unregister chan *Client

	// Broadcast messages to clients
	broadcast chan []byte
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
	// User ID associated with this connection
	userID int
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
