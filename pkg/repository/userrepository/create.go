package userrepository

import (
	"errors"
	"fmt"
	"strings"

	dbconnection "github.com/chaveshigor/todo-api/infra/db"
	"github.com/chaveshigor/todo-api/pkg/model"
)

func Create(new_user *model.User) (*model.User, error) {
	db_instance, _ := dbconnection.Create()
	defer db_instance.Close()

	sql := "INSERT INTO users (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err := db_instance.Connection.Exec(sql, new_user.ID, new_user.Name, new_user.Email, new_user.CreatedAt, new_user.UpdatedAt)
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
