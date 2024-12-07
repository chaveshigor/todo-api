package serializers_test

import (
	"testing"

	"github.com/chaveshigor/todo-api/infra/server/serializers"
	"github.com/chaveshigor/todo-api/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestUserSerializer(t *testing.T) {
	t.Run("Serialize user", func(t *testing.T) {
		user, _ := model.NewUser("higor", "higor@mail.com")

		serialized_user := serializers.SerializeUser(user)

		assert.Equal(t, serialized_user.ID, user.ID.String())
		assert.Equal(t, serialized_user.Email, user.Email)
		assert.Equal(t, serialized_user.Name, user.Name)
	})
}
