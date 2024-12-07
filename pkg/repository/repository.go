package repository

import (
	dbconnection "github.com/chaveshigor/todo-api/infra/db"
	user_repository "github.com/chaveshigor/todo-api/pkg/repository/user-repository"
)

type Repository struct {
	db             *dbconnection.DbInstance
	UserRepository *user_repository.UserRepository
}

func New() (*Repository, error) {
	db, err := dbconnection.Create()
	if err != nil {
		return nil, err
	}

	repository := Repository{
		UserRepository: &user_repository.UserRepository{DB: db},
	}

	return &repository, nil
}

func (r *Repository) CloseDbConnection() {
	r.db.Close()
}
