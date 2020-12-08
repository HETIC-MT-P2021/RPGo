package repository_test

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var character = &repository.Character{
	Name:          "First Character",
	Class:         "ranger",
	DiscordUserID: "1234",
}

type testCaseData struct {
	query string
	rows  *sqlmock.Rows
	char  *repository.Character
}

type testCases map[string]testCaseData

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestCharacterRepository_GetCharacterByDiscordUserID(t *testing.T) {

	cases := testCases{
		"should find existing character": testCaseData{
			query: "SELECT pc.id, pc.name, pc.class, pc.discord_user_id FROM p_character pc WHERE pc.discord_user_id = \\?",
			rows:  sqlmock.NewRows([]string{"id", "name", "class", "discord_user_id"}).AddRow(character.ID, character.Name, character.Class, character.DiscordUserID),
			char:  character,
		},
		"should not find existing character": testCaseData{
			query: "SELECT pc.id, pc.name, pc.class, pc.discord_user_id FROM p_character pc WHERE pc.discord_user_id = \\?",
			rows:  sqlmock.NewRows([]string{"id", "name", "class", "discord_user_id"}),
			char:  nil,
		},
	}

	for index, parameters := range cases {
		t.Run(index, func(t *testing.T) {
			db, mock := NewMock()
			repo := &repository.CharacterRepository{Conn: db}
			defer func() {
				repo.Close()
			}()
			mock.ExpectQuery(parameters.query).WithArgs(character.DiscordUserID).WillReturnRows(parameters.rows)
			character, err := repo.GetCharacterByDiscordUserID(character.DiscordUserID)
			assert.Equal(t, character, parameters.char)
			assert.NoError(t, err)
		})
	}
}

func TestCharacterRepository_Create(t *testing.T) {

	db, mock := NewMock()
	repo := &repository.CharacterRepository{Conn: db}
	defer func() {
		repo.Close()
	}()

	query := "INSERT INTO p_character \\(name, class, discord_user_id\\) VALUES \\(\\?,\\?,\\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(character.Name, character.Class, character.DiscordUserID).WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Create(character)
	assert.NoError(t, err)

	assert.Equal(t, int64(1), character.ID)
}
