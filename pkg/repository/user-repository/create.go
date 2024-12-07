package user_repository

import (
	"errors"
	"fmt"
	"strings"

	"github.com/chaveshigor/todo-api/pkg/model"
)

func (ur UserRepository) Create(new_user *model.User) (*model.User, error) {
	sql := "INSERT INTO users (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"

	_, err := ur.DB.Connection.Exec(sql, new_user.ID, new_user.Name, new_user.Email, new_user.CreatedAt, new_user.UpdatedAt)
	if err != nil {
		fmt.Println(err)
		return nil, adaptError(err)
	}

	return new_user, nil
}

func adaptError(err error) error {
	if strings.Contains(err.Error(), "UNIQUE constraint failed") {
		return errors.New("email already in use")
	}

	return err
}
