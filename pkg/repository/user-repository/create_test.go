package user_repository_test

import (
	"testing"

	"github.com/chaveshigor/todo-api/infra/db/config"
	"github.com/chaveshigor/todo-api/pkg/model"
	"github.com/chaveshigor/todo-api/pkg/repository"
	"github.com/chaveshigor/todo-api/test/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testhelpers.InitializeEnv()

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("valid user", func(t *testing.T) {
		t.Cleanup(cleanDb)

		user, _ := model.NewUser("higor", "higor@mail.com")

		repo, _ := repository.New()
		repo.UserRepository.Create(user)
		user, err := repo.UserRepository.Create(user)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Name, "higor")
		assert.Equal(t, user.Email, "higor@mail.com")
	})

	t.Run("valid user with upper case name", func(t *testing.T) {
		t.Cleanup(cleanDb)

		user, _ := model.NewUser("HIGOR", "HIGOR@mail.com")

		repo, _ := repository.New()
		user, err := repo.UserRepository.Create(user)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, user.Name, "higor")
		assert.Equal(t, user.Email, "higor@mail.com")
	})
}

func cleanDb() {
	config.MigrateDown()
	config.MigrateUp()
}