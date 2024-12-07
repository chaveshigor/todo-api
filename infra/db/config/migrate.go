package config

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateUp() error {
	db_config := getDbConfig()
	migrations, err := migrate.New(
		fmt.Sprintf("file://%s", db_config.MigrationPath),
		fmt.Sprintf("sqlite3://%s", db_config.DbPath),
	)

	if err != nil {
		log.Fatalf("Error initializing migration: %v", err)
	}

	err = migrations.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Println("No need to run migrations.")
			return nil
		} else {
			log.Fatalf("Error running migrations: %v", err)
			return err
		}
	} else {
		fmt.Println("Migrations runned with success!")
		return nil
	}
}

func MigrateDown() error {
	db_config := getDbConfig()
	migrations, err := migrate.New(
		fmt.Sprintf("file://%s", db_config.MigrationPath),
		fmt.Sprintf("sqlite3://%s", db_config.DbPath),
	)

	if err != nil {
		log.Fatalf("Error initializing migration: %v", err)
	}

	err = migrations.Down()
	return err
}
