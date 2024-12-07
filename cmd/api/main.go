package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/chaveshigor/todo-api/infra/server"
	"github.com/chaveshigor/todo-api/pkg/repository"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	repo, err := repository.New()
	if err != nil {
		log.Fatal(err)
	}

	defer repo.CloseDbConnection()

	routes := server.InitializeRoutes(repo)

	port := os.Getenv("APP_PORT")

	fmt.Printf("Server starting at port %s...", port)
	http.Handle("/", routes)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println("Error initializing server:", err)
	}
}
