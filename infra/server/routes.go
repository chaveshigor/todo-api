package server

import (
	"github.com/gorilla/mux"

	"github.com/chaveshigor/todo-api/infra/server/handlers"
)

func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")

	return r
}
