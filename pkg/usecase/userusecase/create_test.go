package userusecase_test

import (
	"testing"

	"github.com/chaveshigor/todo-api/infra/db/config"
	"github.com/chaveshigor/todo-api/pkg/model"
	"github.com/chaveshigor/todo-api/pkg/usecase/userusecase"
	"github.com/chaveshigor/todo-api/test/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testhelpers.InitializeEnv()

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Create user", func(t *testing.T) {
		config.MigrateUp()
		defer config.MigrateDown()

		name := "higor"
		email := "higor@mail.com"

		user, err := userusecase.Create(name, email)

		assert.Nil(t, err, nil)
		assert.IsType(t, &model.User{}, user, "user should be a *User{}")
	})

	t.Run("Return error when email is in use", func(t *testing.T) {
		config.MigrateUp()
		defer config.MigrateDown()

		name := "higor"
		email := "higor@mail.com"

		userusecase.Create(name, email)
		user, err := userusecase.Create(name, email)

		assert.Error(t, err)
		assert.Equal(t, err.Error(), "email already in use")
		assert.Nil(t, user, "user should be nil")
	})
}
