package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	userusecase "github.com/chaveshigor/todo-api/pkg/usecase/user_usecase"
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
		fmt.Println(err)
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	user, err := userusecase.Create(request.User.Name, request.User.Email)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(user)
}
