package repository

import (
	"database/sql"
)

//CharacterRepositoryInterface interface for functions for character repository
type CharacterRepositoryInterface interface {
	Create(character *Character) error
	GetCharacterByDiscordUserID(discordUserID string) (*Character, error)
}

// CharacterRepository struct for db connection
type CharacterRepository struct {
	Conn *sql.DB
}

//Close the DB connection
func (repository *CharacterRepository) Close() {
	repository.Conn.Close()
}

//Character is a playable character model
type Character struct {
	ID            int64
	Name          string
	Class         Class
	DiscordUserID string
	//@toDo add an inventory implem, this is user v0
}

//Class enum type
type Class string

//Defines the possible classes a user can choose
const (
	Rogue  Class = "rogue"
	Knight Class = "knight"
	Wizard Class = "wizard"
)

//IsValid returns true if the class is an existing one
func (class Class) IsValid() bool {
	switch class {
	case Rogue, Knight, Wizard:
		return true
	}
	return false
}
