package user_usecase

import (
	"github.com/chaveshigor/todo-api/pkg/model"
	"github.com/chaveshigor/todo-api/pkg/repository"
)

func Create(repositories *repository.Repository, name, email string) (*model.User, error) {
	newUser, err := model.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	user, err := repositories.UserRepository.Create(newUser)
	if err != nil {
		return nil, err
	}

	return user, nil
}
