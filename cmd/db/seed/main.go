package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/brianvoe/gofakeit/v6"
	database "github.com/devanfer02/ratemyubprof/internal/infra/database/postgres"
	"github.com/devanfer02/ratemyubprof/internal/infra/env"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

var db *sqlx.DB

func main() {
	env := env.NewEnv()
	db = database.NewDatabase(env)

	seedUsers(1000)
	seedReviews(10000)
	seedReviewReactions(90000)
}

func seedUsers(count int) {
	batchSize := 100
	for i := 0; i < count; i += batchSize {
		valueStrings := make([]string, 0, batchSize)
		valueArgs := make([]interface{}, 0, batchSize*4)
		end := i + batchSize
		if end > count {
			end = count
		}

		for j := i; j < end; j++ {
			p1 := len(valueArgs) + 1
			p2 := p1 + 1
			p3 := p1 + 2
			p4 := p1 + 3
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", p1, p2, p3, p4))
			valueArgs = append(valueArgs, gofakeit.UUID(), gofakeit.DigitN(8), gofakeit.Username(), gofakeit.Password(true, true, true, true, false, 12))
		}

		stmt := fmt.Sprintf("INSERT INTO users (id, nim, username, password) VALUES %s", strings.Join(valueStrings, ","))
		_, err := db.Exec(stmt, valueArgs...)
		if err != nil {
			log.Printf("insert user error on batch starting at %d: %v\n", i, err)
		}
		log.Printf("Processed %d users\n", end)
	}
	fmt.Println("Users seeded!")
}

func seedReviews(count int) {
	var userIDs, profIDs []string
	if err := db.Select(&userIDs, "SELECT id FROM users"); err != nil {
		log.Fatalf("error fetching users: %v", err)
	}
	if err := db.Select(&profIDs, "SELECT id FROM professors"); err != nil {
		log.Fatalf("error fetching professors: %v", err)
	}

	if len(userIDs) == 0 || len(profIDs) == 0 {
		log.Println("No users or professors to create reviews for.")
		return
	}

	batchSize := 100
	for i := 0; i < count; i += batchSize {
		valueStrings := make([]string, 0, batchSize)
		valueArgs := make([]interface{}, 0, batchSize*6)
		end := i + batchSize
		if end > count {
			end = count
		}

		usedPairsInBatch := make(map[string]bool)

		for j := i; j < end; j++ {
			userID := userIDs[gofakeit.Number(0, len(userIDs)-1)]
			profID := profIDs[gofakeit.Number(0, len(profIDs)-1)]
			key := userID + "-" + profID

			if usedPairsInBatch[key] {
				continue
			}
			usedPairsInBatch[key] = true
			
			p1 := len(valueArgs) + 1
			p2 := p1 + 1
			p3 := p1 + 2
			p4 := p1 + 3
			p5 := p1 + 4
			p6 := p1 + 5

			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d)", p1, p2, p3, p4, p5, p6))
			valueArgs = append(valueArgs, gofakeit.UUID(), userID, profID, gofakeit.Sentence(10), gofakeit.Float64Range(1, 5), gofakeit.Float64Range(1, 5))
		}

		if len(valueArgs) == 0 {
			continue
		}

		stmt := fmt.Sprintf(`INSERT INTO reviews (id, user_id, prof_id, comment, difficulty_rating, friendly_rating) VALUES %s ON CONFLICT (user_id, prof_id) DO NOTHING`, strings.Join(valueStrings, ","))
		_, err := db.Exec(stmt, valueArgs...)
		if err != nil {
			log.Printf("insert review error on batch starting at %d: %v\n", i, err)
		}
		log.Printf("Processed %d reviews\n", end)
	}

	fmt.Println("Reviews seeded!")
}

func seedReviewReactions(count int) {
	var userIDs, reviewIDs []string
	if err := db.Select(&userIDs, "SELECT id FROM users"); err != nil {
		log.Fatalf("error fetching users: %v", err)
	}
	if err := db.Select(&reviewIDs, "SELECT id FROM reviews"); err != nil {
		log.Fatalf("error fetching reviews: %v", err)
	}

	if len(userIDs) == 0 || len(reviewIDs) == 0 {
		log.Println("No users or reviews to create reactions for.")
		return
	}

	batchSize := 100
	for i := 0; i < count; i += batchSize {
		valueStrings := make([]string, 0, batchSize)
		valueArgs := make([]interface{}, 0, batchSize*3)
		end := i + batchSize
		if end > count {
			end = count
		}

		usedPairsInBatch := make(map[string]bool)

		for j := i; j < end; j++ {
			userID := userIDs[gofakeit.Number(0, len(userIDs)-1)]
			reviewID := reviewIDs[gofakeit.Number(0, len(reviewIDs)-1)]
			key := userID + "-" + reviewID

			if usedPairsInBatch[key] {
				continue
			}
			usedPairsInBatch[key] = true

			p1 := len(valueArgs) + 1
			p2 := p1 + 1
			p3 := p1 + 2
			valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", p1, p2, p3))
			valueArgs = append(valueArgs, userID, reviewID, gofakeit.Number(1, 2))
		}

		if len(valueArgs) == 0 {
			continue
		}

		stmt := fmt.Sprintf(`INSERT INTO review_reactions (user_id, review_id, reaction_type) VALUES %s ON CONFLICT (user_id, review_id) DO NOTHING`, strings.Join(valueStrings, ","))
		_, err := db.Exec(stmt, valueArgs...)
		if err != nil {
			log.Printf("insert review_reaction error on batch starting at %d: %v\n", i, err)
		}
		log.Printf("Processed %d review reactions\n", end)
	}

	fmt.Println("Review reactions seeded!")
}