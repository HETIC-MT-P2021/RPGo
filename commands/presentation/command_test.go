package presentation_test

import (
	"fmt"
	"github.com/HETIC-MT-P2021/RPGo/commands"
	"github.com/HETIC-MT-P2021/RPGo/commands/presentation"
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

type testCaseData struct {
	character *repository.Character
	answer    string
}

type testCases map[string]testCaseData

func TestCharCommandGenerator_Presentation(t *testing.T) {
	cases := testCases{
		"character does exist": {
			character: char,
			answer:    fmt.Sprintf(helpers.CharacterPresentation, char.Name, char.Class),
		},
		"character does not exist": {
			character: nil,
			answer:    helpers.CharacterDoesNotExist,
		},
	}
	for index, parameters := range cases {
		t.Run(index, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			m := repository.NewMockCharacterRepositoryInterface(ctrl)

			m.EXPECT().GetCharacterByDiscordUserID(char.DiscordUserID).DoAndReturn(func(_ string) (*repository.
				Character, error) {
				return parameters.character, nil
			}).Times(1)

			s := commands.NewMockDiscordConnector(ctrl)

			message := discordgo.Message{
				ID:        "428",
				ChannelID: "43",
				GuildID:   "",
				Content:   "character presentation test",
			}

			messageCreate := discordgo.MessageCreate{Message: &message}

			generator := presentation.CharCommandGenerator{Repo: m}

			charPresentationCommand, err := generator.PresentationCommand(s, &messageCreate, char.DiscordUserID)
			require.NoError(t, err, "should create presentation command")
			assert.Equal(t, parameters.answer, charPresentationCommand.Payload().Answer)
		})
	}
}
