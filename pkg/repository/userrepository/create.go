package userrepository

import (
	"log"

	dbconnection "github.com/chaveshigor/todo-api/infra/db"
	"github.com/chaveshigor/todo-api/pkg/model"
)

func Create(db_instance *dbconnection.DbInstance, name, email string) (*model.User, error) {
	new_user, err := model.NewUser(name, email)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	sql := "INSERT INTO users (id, name, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	_, err = db_instance.Connection.Exec(sql, new_user.ID, new_user.Name, new_user.Email, new_user.CreatedAt, new_user.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return new_user, nil
}
