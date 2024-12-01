package dbconnection

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type DbInstance struct {
	Connection *sql.DB
}

func Create() (*DbInstance, error) {
	db, err := sql.Open("sqlite3", database_path())
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	Db_instance := DbInstance{Connection: db}
	return &Db_instance, nil
}

func (i *DbInstance) Close() {
	i.Connection.Close()
}

func database_path() string {
	base_path := "/home/higor/projects/personal/todo-api/infra/db/"
	current_env := os.Getenv("APP_ENV")

	return fmt.Sprintf("%s%s.db", base_path, current_env)
}
