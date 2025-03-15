package main

import (
	"database/sql"
	"fmt"
	"github.com/hako/branca"
	"github.com/stymsinghss/Tweety/internal/handler"
	"github.com/stymsinghss/Tweety/internal/service"
	"log"
	"net/http"
)

const (
	databaseURL = "postgresql://root@127.0.0.1:26257/tweety?sslmode=disable"
	port = 3000
)

func main() {
	fmt.Println("Tweety")

	// Connect to database
	database, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database. Failed with -> %v\n", err)
		return
	}
	defer database.Close()

	// Ping database
	if err = database.Ping(); err != nil {
		log.Fatalf("Error pinging database. Failed with -> %v\n", err)
		return
	}

	// Setup Branca token
	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit")

	// Setup service
	svc := service.New(database, codec)

	// Create handlers and pass service to it
	h := handler.New(svc)

	// Create server
	addr := fmt.Sprintf(":%d", port)
	if err = http.ListenAndServe(addr, h); err != nil {
		log.Fatalf("❌ Server failed: %v\n", err)
	}
	log.Printf("🚀 Server is running on http://localhost:%d", port)
}
