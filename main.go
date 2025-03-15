package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/hako/branca"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/stymsinghss/Tweety/internal/handler"
	"github.com/stymsinghss/Tweety/internal/service"
)

const (
	databaseURL = "postgresql://root@127.0.0.1:26257/tweety?sslmode=disable"
	port        = 3000
	secretKey   = "supersecretkeyyoushouldnotcommit"
)

func main() {
	fmt.Println("🚀 Tweety is starting...")

	// Initialize database connection
	db := mustInitDB()
	defer db.Close()

	// Setup Branca token
	codec := branca.NewBranca(secretKey)

	// Setup service & handler
	svc := service.New(db, codec)
	h := handler.New(svc)

	// Start HTTP server
	startServer(h)
}

// mustInitDB initializes and verifies the database connection.
func mustInitDB() *sql.DB {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("❌ Failed to ping database: %v", err)
	}

	log.Println("✅ Connected to database successfully!")
	return db
}

// startServer starts the HTTP server.
func startServer(h http.Handler) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("🚀 Server is running at http://localhost%s", addr)

	if err := http.ListenAndServe(addr, h); err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
