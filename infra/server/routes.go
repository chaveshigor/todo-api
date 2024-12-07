package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/chaveshigor/todo-api/infra/server/handlers"
	"github.com/chaveshigor/todo-api/pkg/repository"
)

func InitializeRoutes(repo *repository.Repository) *mux.Router {
	r := mux.NewRouter()

	useRoute(r, "POST", "/users", repo, handlers.CreateUserHandler)

	return r
}

func useRoute(
	r *mux.Router,
	method, path string,
	repo *repository.Repository,
	handler func(w http.ResponseWriter, r *http.Request, repo *repository.Repository),
) {
	r.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, repo)
	}).Methods(method)
}
