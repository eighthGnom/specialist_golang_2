package storage

import (
	"fmt"
	"log"

	"github.com/eighthGnom/standard_web_server/internal/app/models"
)

type UserRepository struct {
	storage *Storage
}

var (
	userTable = "users"
)

func (ur *UserRepository) Create(user *models.User) (*models.User, error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password) VALUES ($1, $2) RETURNING user_id", userTable)
	if err := ur.storage.db.QueryRow(query, user.Login, user.Password).Scan(&user.ID); err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *UserRepository) FindByLogin(login string) (*models.User, bool, error) {
	users, err := ur.SelectAll()
	var founded bool
	if err != nil {
		return nil, founded, err
	}
	var foundedUser *models.User
	for _, user := range users {
		if user.Login == login {
			foundedUser = user
			founded = true
			break
		}
	}
	return foundedUser, founded, nil
}

func (ur *UserRepository) SelectAll() ([]*models.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s", userTable)
	rows, err := ur.storage.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*models.User, 0)
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.ID, &user.Login, &user.Password); err != nil {
			log.Println(err)
			continue
		}
		users = append(users, &user)
	}
	return users, nil
}
