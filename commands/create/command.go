package create

import (
	"fmt"
	"github.com/HETIC-MT-P2021/RPGo/commands"
	"github.com/HETIC-MT-P2021/RPGo/helpers"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/bwmarrin/discordgo"
)

//CharacterCreateCommand model for character creation command
type CharacterCreateCommand struct {
	Receiver *Receiver
	payload  *CharacterCreateCommandPayload
}

//CharacterCreateCommandPayload model for information on character creation command
type CharacterCreateCommandPayload struct {
	Answer  string
	session commands.DiscordConnectorMessage
	Message *discordgo.MessageCreate
}

//CharCommandGenerator stores the repository.CharacterRepository DTO
type CharCommandGenerator struct {
	Repo repository.CharacterRepositoryInterface
}

//CreateCommand a character creation command
func (command *CharCommandGenerator) CreateCommand(c commands.DiscordConnectorMessage, m *discordgo.MessageCreate,
	name string, class commands.Class, userID string) (*CharacterCreateCommand, error) {

	char, err := command.Repo.GetCharacterByDiscordUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("couldn't get character with userID: %s, %v", userID, err)
	}

	if !class.IsValid() {

		return &CharacterCreateCommand{
			Receiver: &Receiver{},
			payload: &CharacterCreateCommandPayload{
				Answer:  helpers.WrongClassGiven,
				session: c,
				Message: m,
			}}, nil
	}

	if char != nil {

		return &CharacterCreateCommand{
			Receiver: &Receiver{},
			payload: &CharacterCreateCommandPayload{
				Answer:  helpers.CharAlreadyExists,
				session: c,
				Message: m,
			}}, nil
	}

	//@toDo add newCreateCharacterCommand

	character := repository.Character{
		Name:          name,
		Class:         class,
		DiscordUserID: userID,
	}

	err = command.Repo.Create(&character)
	if err != nil {
		return nil, fmt.Errorf("couldn't create character: %v", err)
	}

	return &CharacterCreateCommand{
		Receiver: &Receiver{},
		payload: &CharacterCreateCommandPayload{
			Answer:  fmt.Sprintf(helpers.CharSuccessfullyCreated, name),
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
func (p *CharacterCreateCommandPayload) Session() commands.DiscordConnectorMessage {
	return p.session
}
