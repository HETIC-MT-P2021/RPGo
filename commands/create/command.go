package create

import (
	"fmt"
	"github.com/HETIC-MT-P2021/RPGo/database"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	_ "github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/bwmarrin/discordgo"
	"log"
)

type CreateCommand struct {
	Receiver *Receiver
	Payload  *CreateCommandPayload
}

type CreateCommandPayload struct {
	Answer  string
	Session *discordgo.Session
	Message *discordgo.MessageCreate
}

func MakeCreateCommand(s *discordgo.Session, m *discordgo.MessageCreate, name string, userID string) *CreateCommand {
	db := database.DBCon
	repo := repository.Repository{Conn: db}
	answer := "You already have a character!"

	char, err := repo.GetCharacterByDiscordUserID(userID)
	if err != nil {
		log.Printf("Couldn't get character with userID: %s, %v", userID, err)
	}

	// Create a character if none found in DB
	if char == nil {
		character := repository.Character{
			Name:          name,
			Class:         "Ranger",
			DiscordUserID: userID,
		}
		repo.CreateACharacter(&character)
		answer = fmt.Sprintf("%s successfully created!", name)
	}

	return &CreateCommand{
		Receiver: &Receiver{},
		Payload: &CreateCommandPayload{
			Answer:  answer,
			Session: s,
			Message: m,
		},
	}
}

func (c *CreateCommand) Execute() {
	c.Receiver.Answer(c.Payload)
}
