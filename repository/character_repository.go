package repository

import "database/sql"

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
	ID              int64
	Name            string
	Class           string
	DiscordUserID   string
	DiscordServerID int64 // @toDo(team) : decide if user only gets one PC or one per server
	//@toDo add an inventory implem, this is user v0
}
