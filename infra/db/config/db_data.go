package config

import (
	"fmt"
	"os"
)

type DbConfig struct {
	DbPath        string
	MigrationPath string
}

func getDbConfig() *DbConfig {
	return &DbConfig{
		DbPath:        db_path(),
		MigrationPath: migrations_path(),
	}
}

func db_path() string {
	base_path := "/home/higor/projects/personal/todo-api/infra/db/"
	current_env := os.Getenv("APP_ENV")

	return fmt.Sprintf("%s%s.db", base_path, current_env)
}

func migrations_path() string {
	return "/home/higor/projects/personal/todo-api/infra/db/migrations"
}
