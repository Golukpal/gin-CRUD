package db

import (
	"log"
	"os"
)

func RunMigrations() {
	files := []string{
		"migrations/001_create_users.sql",
		"migrations/002_create_posts.sql",
	}

	for _, file := range files {
		sqlBytes, err := os.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}

		_, err = DB.Exec(string(sqlBytes))
		if err != nil {
			log.Fatal(err)
		}
	}
}