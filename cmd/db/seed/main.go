package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	database "github.com/devanfer02/ratemyubprof/internal/infra/database/postgres"
	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {
	var err error
	env := env.NewEnv()
	db = database.NewDatabase(env)
	if err != nil {
		log.Fatalln(err)
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go func(){
		defer wg.Done()
		seedUsers(50000)
	}()
	go func(){
		defer wg.Done()
		seedReviews(100000)
	}()
	go func(){
		defer wg.Done()
		seedReviewReactions(100000)
	}()
	

	wg.Wait()
	
}

func seedUsers(count int) {
	for i := 0; i < count; i++ {
		id := gofakeit.UUID()
		nim := gofakeit.DigitN(8)
		username := gofakeit.Username()
		password := gofakeit.Password(true, true, true, true, false, 12)

		_, err := db.Exec(`INSERT INTO users (id, nim, username, password) 
			VALUES ($1, $2, $3, $4)`,
			id, nim, username, password)
		if err != nil {
			log.Println("insert user error:", err)
		}

		log.Println("Inserted user:", i+1)
	}
	fmt.Println("Users seeded!")
}

func seedReviews(count int) {
	var userIDs, profIDs []string
	_ = db.Select(&userIDs, "SELECT id FROM users")
	_ = db.Select(&profIDs, "SELECT id FROM professors")

	var wg sync.WaitGroup
	var mu sync.Mutex
	usedPairs := make(map[string]bool)

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			userID := userIDs[gofakeit.Number(0, len(userIDs)-1)]
			profID := profIDs[gofakeit.Number(0, len(profIDs)-1)]
			key := userID + "-" + profID

			mu.Lock()
			if usedPairs[key] {
				mu.Unlock()
				return
			}
			usedPairs[key] = true
			mu.Unlock()

			comment := gofakeit.Sentence(10)
			diffRating := gofakeit.Float64Range(1, 5)
			friendRating := gofakeit.Float64Range(1, 5)

			_, err := db.Exec(`INSERT INTO reviews (id, user_id, prof_id, comment, difficulty_rating, friendly_rating, created_at)
				VALUES ($1, $2, $3, $4, $5, $6, $7)
				ON CONFLICT DO NOTHING`,
				gofakeit.UUID(), userID, profID, comment, diffRating, friendRating, time.Now())

			if err != nil {
				log.Println("insert review error:", err)
			}

			log.Println("Inserted review:", i+1)
		}()
	}

	wg.Wait()
	fmt.Println("Reviews seeded concurrently!")
}

func seedReviewReactions(count int) {
	var userIDs, reviewIDs []string
	err := db.Select(&userIDs, "SELECT id FROM users")
	if err != nil {
		log.Println("error fetching users:", err)
		return
	}
	err = db.Select(&reviewIDs, "SELECT id FROM reviews")
	if err != nil {
		log.Println("error fetching reviews:", err)
		return
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	usedPairs := make(map[string]bool)

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			userID := userIDs[gofakeit.Number(0, len(userIDs)-1)]
			reviewID := reviewIDs[gofakeit.Number(0, len(reviewIDs)-1)]
			key := userID + "-" + reviewID

			mu.Lock()
			if usedPairs[key] {
				mu.Unlock()
				return
			}
			usedPairs[key] = true
			mu.Unlock()

			reactionType := gofakeit.Number(1, 2)

			_, err := db.Exec(`
				INSERT INTO review_reactions (user_id, review_id, reaction_type)
				VALUES ($1, $2, $3)
				ON CONFLICT DO NOTHING`,
				userID, reviewID, reactionType)

			if err != nil {
				log.Println("insert review_reaction error:", err)
			}
			log.Println("Inserted review reaction:", i+1)
		}()
	}

	wg.Wait()
	fmt.Println("Review reactions seeded concurrently!")
}
