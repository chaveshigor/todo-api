package userusecase

import (
	dbconnection "github.com/chaveshigor/todo-api/infra/db"
	"github.com/chaveshigor/todo-api/pkg/model"
	"github.com/chaveshigor/todo-api/pkg/repository/userrepository"
)

func Create(name, email string) (*model.User, error) {
	db, _ := dbconnection.Create()

	user, err := userrepository.Create(db, name, email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
