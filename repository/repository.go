package repository

import "database/sql"

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
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
