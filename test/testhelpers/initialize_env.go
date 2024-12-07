package testhelpers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitializeEnv() {
	err := godotenv.Load("/home/higor/projects/personal/todo-api/.env")
	if err != nil {
		fmt.Println(err)
	}

	err = os.Setenv("APP_ENV", "test")
	if err != nil {
		log.Fatal(err, "Error subscribing env variable APP_ENV")
	}
}
