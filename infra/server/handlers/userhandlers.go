package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chaveshigor/todo-api/infra/server/serializers"
	"github.com/chaveshigor/todo-api/pkg/usecase/userusecase"
)

type request struct {
	User struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var request request

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	user, err := userusecase.Create(request.User.Name, request.User.Email)
	if err != nil {
		serializedError, _ := json.Marshal(serializers.SerializeError(err))
		http.Error(w, string(serializedError), http.StatusBadRequest)
		return
	}

	serializedUser := serializers.SerializeUser(user)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(serializedUser)
}
