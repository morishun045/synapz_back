// main.go

package main

import (
	"log"
	"synapz_back/db"
	"synapz_back/model"
)

func main() {
	db.Init()

	if db.DB != nil {
		err := db.DB.AutoMigrate(
			&model.User{},
			&model.Post{},
			&model.Reply{},
			&model.Like{},
			&model.Retweet{},
			&model.Article{},
			&model.Bookmark{},
			&model.Follow{},
		)

		if err != nil {
			log.Fatalf("Migration Failed: %v", err)
		}

		log.Println("Migration Completely success")
	} else {
		log.Println("Skipping migration - no database connection")
	}
}
