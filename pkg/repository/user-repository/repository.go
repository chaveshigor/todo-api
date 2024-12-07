package user_repository

import dbconnection "github.com/chaveshigor/todo-api/infra/db"

type UserRepository struct {
	DB *dbconnection.DbInstance
}
