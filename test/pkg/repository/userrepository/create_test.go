package userrepository_test

import (
	"log"
	"os"
	"testing"

	dbconnection "github.com/chaveshigor/todo-api/infra/db"
	"github.com/chaveshigor/todo-api/infra/db/config"
	"github.com/chaveshigor/todo-api/pkg/repository/userrepository"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("/home/higor/projects/personal/todo-api/.env")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Setenv("APP_ENV", "test")
	if err != nil {
		log.Fatal(err, "Error subscribing env variable APP_ENV")
	}

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("valid user", func(t *testing.T) {
		config.MigrateUp()
		defer config.MigrateDown()

		db_instance, err := dbconnection.Create()

		if err != nil {
			t.Error(err)
		}

		defer db_instance.Close()

		user, err := userrepository.Create(db_instance, "higor", "higor@mail.com")
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Name, "higor")
		assert.Equal(t, user.Email, "higor@mail.com")
	})

	t.Run("valid user with upper case name", func(t *testing.T) {
		config.MigrateUp()
		defer config.MigrateDown()

		db_instance, err := dbconnection.Create()

		if err != nil {
			t.Error(err)
		}

		defer db_instance.Close()

		user, err := userrepository.Create(db_instance, "HIGOR", "higor@mail.com")
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Name, "higor")
		assert.Equal(t, user.Email, "higor@mail.com")
	})
}
