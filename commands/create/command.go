package create

import (
	"fmt"
	"github.com/HETIC-MT-P2021/RPGo/database"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/bwmarrin/discordgo"
	"log"
)

type CharacterCreateCommand struct {
	Receiver *Receiver
	payload  *CharacterCreateCommandPayload
}

type CharacterCreateCommandPayload struct {
	Answer  string
	session *discordgo.Session
	Message *discordgo.MessageCreate
}

func MakeCreateCommand(s *discordgo.Session, m *discordgo.MessageCreate, name string, userID string) *CharacterCreateCommand {
	db := database.DBCon
	repo := repository.Repository{Conn: db}
	answer := "You already have a character!"

	char, err := repo.GetCharacterByDiscordUserID(userID)
	if err != nil {
		log.Fatalf("Couldn't get character with userID: %s, %v", userID, err)
	}

	// Create a character if none found in DB
	if char == nil {
		character := repository.Character{
			Name:          name,
			Class:         "Ranger",
			DiscordUserID: userID,
		}
		err := repo.CreateACharacter(&character)
		if err != nil {
			log.Fatalf("Couldn't create character: %v", err)
		}
		answer = fmt.Sprintf("%s successfully created!", name)
	}

	return &CharacterCreateCommand{
		Receiver: &Receiver{},
		payload: &CharacterCreateCommandPayload{
			Answer:  answer,
			session: s,
			Message: m,
		},
	}
}

func (c *CharacterCreateCommand) Execute() {
	c.Receiver.Answer(c.Payload())
}
func (c *CharacterCreateCommand) Payload() *CharacterCreateCommandPayload {
	return c.payload
}

func (p *CharacterCreateCommandPayload) Session() *discordgo.Session {
	return p.session
}
