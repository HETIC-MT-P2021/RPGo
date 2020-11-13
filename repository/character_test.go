package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var c = &Character{
	Name:          "First Character",
	Class:         "ranger",
	DiscordUserID: "1234",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestRepository_CreateACharacter(t *testing.T) {

	db, mock := NewMock()
	repo := &CharacterRepository{db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO p_character \\(name, class, discord_user_id\\) VALUES \\(\\?,\\?,\\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(c.Name, c.Class, c.DiscordUserID).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(c)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), c.ID)
}
