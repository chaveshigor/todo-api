package main

import (
	"fmt"

	"github.com/chaveshigor/todo-api/infra/db/config"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
	}

	config.MigrateUp()
}
