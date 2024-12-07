package handlers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chaveshigor/todo-api/infra/db/config"
	"github.com/chaveshigor/todo-api/infra/server/handlers"
	"github.com/chaveshigor/todo-api/test/testhelpers"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	testhelpers.InitializeEnv()

	m.Run()
}

func TestUserCreation(t *testing.T) {
	t.Run("valid user", func(t *testing.T) {
		t.Cleanup(cleanDb)

		reqBody := `{"user": {"name": "higor", "email": "higor@mail.com"}}`
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer([]byte(reqBody)))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.CreateUserHandler)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
	})

	t.Run("email in use", func(t *testing.T) {
		t.Cleanup(cleanDb)

	})

	t.Run("blank email", func(t *testing.T) {
		t.Cleanup(cleanDb)

	})

	t.Run("blank name", func(t *testing.T) {
		t.Cleanup(cleanDb)

	})
}

func cleanDb() {
	config.MigrateDown()
	config.MigrateUp()
}
