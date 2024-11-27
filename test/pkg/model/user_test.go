package model_test

import (
	"testing"

	"github.com/chaveshigor/todo-api/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	t.Run("valid user", func(t *testing.T) {
		user, _ := model.NewUser("john", "john@example.com")

		assert.NotEmpty(t, user.ID, "ID should not be empty")
		assert.Equal(t, "john", user.Name, "Name should be 'john'")
		assert.Equal(t, "john@example.com", user.Email, "Email should be 'john@example.com'")
	})

	t.Run("name is required", func(t *testing.T) {
		_, err := model.NewUser("", "mail@test.com")

		assert.Error(t, err)
	})

	t.Run("email is required", func(t *testing.T) {
		_, err := model.NewUser("name", "")

		assert.Error(t, err)
	})
}
