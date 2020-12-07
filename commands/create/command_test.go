package create_test

import (
	"fmt"
	"github.com/HETIC-MT-P2021/RPGo/commands"
	"github.com/HETIC-MT-P2021/RPGo/commands/create"
	"github.com/HETIC-MT-P2021/RPGo/helpers"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var char = &repository.Character{
	Name:          "TestChar",
	Class:         "rogue",
	DiscordUserID: "1234",
}

func TestCharCommandGenerator_Create(t *testing.T) {
	for index, parameters := range map[string]struct {
		character *repository.Character
		answer    string
		class     repository.Class
	}{
		"char does not exist": {
			character: nil,
			answer:    fmt.Sprintf(helpers.CharSuccessfullyCreated, char.Name),
			class:     char.Class,
		},
		"character does exist": {
			character: char,
			answer:    helpers.CharAlreadyExists,
			class:     char.Class,
		},
		"wrong class given": {
			character: nil,
			answer:    helpers.WrongClassGiven,
			class:     "paladin",
		},
	} {
		t.Run(index, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := repository.NewMockCharacterRepositoryInterface(ctrl)

			m.EXPECT().GetCharacterByDiscordUserID(char.DiscordUserID).DoAndReturn(func(_ string) (*repository.
				Character, error) {
				return parameters.character, nil
			}).MaxTimes(1)

			if index != "wrong_class_given" {
				// We don't call the create function as it will not be called if the class is not valid
				m.EXPECT().Create(char).DoAndReturn(func(character *repository.Character) error {
					return nil
				}).MaxTimes(1)
			}

			s := commands.NewMockDiscordConnector(ctrl)

			message := discordgo.Message{
				ID:        "428",
				ChannelID: "43",
				GuildID:   "",
				Content:   "character creation test",
			}

			messageCreate := discordgo.MessageCreate{Message: &message}

			generator := create.CharCommandGenerator{Repo: m}

			charCreateCommand, err := generator.CreateCommand(s, &messageCreate, char.Name, parameters.class, char.DiscordUserID)
			require.NoError(t, err, "should create char command")
			assert.Equal(t, parameters.answer,
				charCreateCommand.Payload().Answer)
		})
	}
}
