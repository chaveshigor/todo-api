package main

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3" // driver SQLite3
	_ "github.com/golang-migrate/migrate/v4/source/file"      // driver de migração (arquivo)
)

func main() {
	migrationsPath := "file://./infra/db/migrations"
	dbURL := "sqlite3://./infra/db/development.db"

	migrations, err := migrate.New(migrationsPath, dbURL)
	if err != nil {
		log.Fatalf("Error initializing migration: %v", err)
	}

	err = migrations.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("No need to run migrations.")
		} else {
			log.Fatalf("Error running migrations: %v", err)
		}
	} else {
		fmt.Println("Migrations runned with success!")
	}
}
