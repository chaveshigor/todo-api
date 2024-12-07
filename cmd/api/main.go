package main

import (
	"fmt"
	"net/http"

	"github.com/chaveshigor/todo-api/infra/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	routes := server.InitializeRoutes()

	// port := 

	fmt.Println("Server starting at port 8080...")
	http.Handle("/", routes)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error initializing server:", err)
	}
}
