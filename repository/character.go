package repository

import "database/sql"

//GetCharacterByDiscordUserID get a character by its discord user id (unique)
func (repository *CharacterRepository) GetCharacterByDiscordUserID(discordUserID string) (*Character, error) {
	row := repository.Conn.QueryRow(`
		SELECT 
		pc.id, 
		pc.name, 
		pc.class, 
		pc.discord_user_id 
		FROM p_character pc 
		WHERE pc.discord_user_id = ?`, discordUserID)
	char := Character{}
	switch err := row.Scan(&char.ID, &char.Name, &char.Class, &char.DiscordUserID); err {
	case sql.ErrNoRows:
		return nil, nil
	case nil:
		return &char, nil
	default:
		return nil, err
	}
}

//Create saves a character in db
func (repository *CharacterRepository) Create(character *Character) error {
	stmt, err := repository.Conn.Prepare(`
	INSERT INTO p_character (name, class, discord_user_id) 
	VALUES (?,?,?)`)
	if err != nil {
		return err
	}
	res, errExec := stmt.Exec(character.Name, character.Class,
		character.DiscordUserID)
	if errExec != nil {
		return errExec
	}

	lastInsertedID, errInsert := res.LastInsertId()
	if errInsert != nil {
		return errInsert
	}

	character.ID = lastInsertedID

	return nil
}
