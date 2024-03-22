package database

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	q := &Queries{Db: db}

	// Prepare the expected result
	expectedUser := User{
		ID:        uuid.New(),
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
		Name:      "Test User",
	}

	// Expect a query to be executed
	mock.ExpectQuery("INSERT INTO users").
		WithArgs(expectedUser.ID, expectedUser.CreatedAt, expectedUser.UpdatedAt, expectedUser.Name).
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name"}).
			AddRow(expectedUser.ID, expectedUser.CreatedAt, expectedUser.UpdatedAt, expectedUser.Name))

	// Call the CreateUser function
	actualUser, err := q.CreateUser(context.Background(), CreateUserParams{
		ID:        expectedUser.ID,
		CreatedAt: expectedUser.CreatedAt,
		UpdatedAt: expectedUser.UpdatedAt,
		Name:      expectedUser.Name,
	})

	// Assert that there was no error and the result is as expected
	require.NoError(t, err)
	require.Equal(t, expectedUser, actualUser)

	// Assert that all expectations were met
	require.NoError(t, mock.ExpectationsWereMet())
}
