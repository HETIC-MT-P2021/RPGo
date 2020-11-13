package create

import (
	"fmt"
	"github.com/HETIC-MT-P2021/RPGo/commands"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/bwmarrin/discordgo"
)

// Messages sent to discord API
const (
	CharAlreadyExists       = "You already have a character!"
	CharSuccessfullyCreated = "%s successfully created!"
)

//CharacterCreateCommand model for character creation command
type CharacterCreateCommand struct {
	Receiver *Receiver
	payload  *CharacterCreateCommandPayload
}

//CharacterCreateCommandPayload model for information on character creation command
type CharacterCreateCommandPayload struct {
	Answer  string
	session commands.DiscordConnector
	Message *discordgo.MessageCreate
}

//CharCommandGenerator stores the repository.CharacterRepository DTO
type CharCommandGenerator struct {
	Repo repository.CharacterRepositoryInterface
}

//Create a character creation command
func (command *CharCommandGenerator) Create(c commands.DiscordConnector, m *discordgo.MessageCreate,
	name string, userID string) (*CharacterCreateCommand, error) {

	answer := CharAlreadyExists
	char, err := command.Repo.GetCharacterByDiscordUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get character with userID: %s, %v", userID, err)
	}

	// Create a character if none found in DB
	if char == nil {
		character := repository.Character{
			Name:          name,
			Class:         "Ranger",
			DiscordUserID: userID,
		}
		err := command.Repo.Create(&character)
		if err != nil {
			return nil, fmt.Errorf("couldn't create character: %v", err)
		}
		answer = fmt.Sprintf(CharSuccessfullyCreated, name)
	}

	return &CharacterCreateCommand{
		Receiver: &Receiver{},
		payload: &CharacterCreateCommandPayload{
			Answer:  answer,
			session: c,
			Message: m,
		},
	}, nil
}

//Execute command with all its information
func (c *CharacterCreateCommand) Execute() {
	c.Receiver.Answer(c.Payload())
}

//Payload returns CharacterCreateCommand payload
func (c *CharacterCreateCommand) Payload() *CharacterCreateCommandPayload {
	return c.payload
}

//Session returns CharacterCreateCommand session
func (p *CharacterCreateCommandPayload) Session() commands.DiscordConnector {
	return p.session
}
