package presentation

import (
	"fmt"
	"github.com/HETIC-MT-P2021/RPGo/commands"
	"github.com/HETIC-MT-P2021/RPGo/helpers"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/bwmarrin/discordgo"
)

//CharacterPresentationCommand model for user's character presentation
type CharacterPresentationCommand struct {
	Receiver *Receiver
	payload  *CharacterPresentationCommandPayload
}

//CharacterPresentationCommandPayload model for information on character presentation command
type CharacterPresentationCommandPayload struct {
	Answer  string
	session commands.DiscordConnector
	Message *discordgo.MessageCreate
}

//CharCommandGenerator stores the repository.CharacterRepository DTO
type CharCommandGenerator struct {
	Repo repository.CharacterRepositoryInterface
}

//PresentationCommand a character creation command
func (command *CharCommandGenerator) PresentationCommand(c commands.DiscordConnector, m *discordgo.MessageCreate,
	userID string) (*CharacterPresentationCommand, error) {

	var answer string

	char, err := command.Repo.GetCharacterByDiscordUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get character with userID: %s, %v", userID, err)
	}

	// CreateCommand a character if none found in DB
	if char == nil {
		answer = helpers.CharacterDoesNotExist
	} else {
		answer = fmt.Sprintf(helpers.CharacterPresentation, char.Name, char.Class)
	}

	return &CharacterPresentationCommand{
		Receiver: &Receiver{},
		payload: &CharacterPresentationCommandPayload{
			Answer:  answer,
			session: c,
			Message: m,
		},
	}, nil
}

//Execute command with all its information
func (c *CharacterPresentationCommand) Execute() {
	c.Receiver.Answer(c.Payload())
}

//Payload returns CharacterCreateCommand payload
func (c *CharacterPresentationCommand) Payload() *CharacterPresentationCommandPayload {
	return c.payload
}

//Session returns CharacterCreateCommand session
func (p *CharacterPresentationCommandPayload) Session() commands.DiscordConnector {
	return p.session
}
