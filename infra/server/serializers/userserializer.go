package serializers

import (
	"github.com/chaveshigor/todo-api/pkg/model"
)

type SerializedUser struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SerializeUser(user *model.User) SerializedUser {
	return SerializedUser{
		ID:    user.ID.String(),
		Name:  user.Name,
		Email: user.Email,
	}
}
