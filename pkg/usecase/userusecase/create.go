package userusecase

import (
	"github.com/chaveshigor/todo-api/pkg/model"
	"github.com/chaveshigor/todo-api/pkg/repository/userrepository"
)

func Create(name, email string) (*model.User, error) {
	new_user, err := model.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	user, err := userrepository.Create(new_user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
