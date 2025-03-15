package database

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// User represents the request payload for registration
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GenerateRandomUsers creates a list of unique dummy users
func GenerateRandomUsers(count int) []User {
	rand.Seed(time.Now().UnixNano()) // Seed random generator

	users := make([]User, count)
	for i := 0; i < count; i++ {
		username := fmt.Sprintf("user_%d", i+1)
		email := fmt.Sprintf("user%d@example.com", i+1)
		password := "password123" // Keep the password same for all users

		users[i] = User{Username: username, Email: email, Password: password}
	}

	return users
}

// Seed calls the user registration handler to insert dummy users
func Seed() {
	log.Println("🌱 Seeding database via API requests...")

	users := GenerateRandomUsers(10) // Generate 40 users

	for _, user := range users {
		payload, _ := json.Marshal(user)
		resp, err := http.Post("http://localhost:3000/api/register", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			log.Printf("❌ Error seeding user %s: %v\n", user.Username, err)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			log.Printf("⚠️ User %s could not be created, status: %d\n", user.Username, resp.StatusCode)
		} else {
			log.Printf("✅ User %s seeded successfully", user.Username)
		}
	}
}