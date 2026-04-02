package repository

import (
	"github.com/Golukpal/gin-crud/internal/db"
	"github.com/Golukpal/gin-crud/internal/models"
)

func CreateUser(user models.User) error {
	_, err := db.DB.Exec(
		"INSERT INTO users(name, email) VALUES($1, $2)",
		user.Name, user.Email,
	)
	return err
}

func GetUsers() ([]models.User, error) {
	rows, err := db.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Name, &u.Email)
		users = append(users, u)
	}
	return users, nil
}