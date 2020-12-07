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

func TestCharCommandGenerator_Create_UserDoesNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repository.NewMockCharacterRepositoryInterface(ctrl)

	m.EXPECT().GetCharacterByDiscordUserID(char.DiscordUserID).DoAndReturn(func(_ string) (*repository.
		Character, error) {
		return nil, nil
	}).MaxTimes(1)

	m.EXPECT().Create(char).DoAndReturn(func(character *repository.Character) error {
		return nil
	}).MaxTimes(1)

	s := commands.NewMockDiscordConnector(ctrl)

	message := discordgo.Message{
		ID:        "428",
		ChannelID: "43",
		GuildID:   "",
		Content:   "Test content",
	}

	messageCreate := discordgo.MessageCreate{Message: &message}

	generator := create.CharCommandGenerator{Repo: m}

	charCreateCommand, err := generator.CreateCommand(s, &messageCreate, "TestChar", "rogue", "1234")
	require.NoError(t, err, "should create char command")
	assert.Equal(t, fmt.Sprintf(helpers.CharSuccessfullyCreated, char.Name),
		charCreateCommand.Payload.Answer)

}

func TestCharCommandGenerator_Create_UserExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repository.NewMockCharacterRepositoryInterface(ctrl)

	m.EXPECT().GetCharacterByDiscordUserID(char.DiscordUserID).DoAndReturn(func(_ string) (*repository.
		Character, error) {
		return char, nil
	}).MaxTimes(1)

	s := commands.NewMockDiscordConnector(ctrl)

	message := discordgo.Message{
		ID:        "428",
		ChannelID: "43",
		GuildID:   "",
		Content:   "Test content",
	}

	messageCreate := discordgo.MessageCreate{Message: &message}

	generator := create.CharCommandGenerator{Repo: m}

	charCreateCommand, err := generator.CreateCommand(s, &messageCreate, "TestChar", "rogue", "1234")
	require.NoError(t, err, "should create char command")
	assert.Equal(t, helpers.CharAlreadyExists, charCreateCommand.Payload.Answer)
}

func TestCharCommandGenerator_Create_WrongClassGiven(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := repository.NewMockCharacterRepositoryInterface(ctrl)

	m.EXPECT().GetCharacterByDiscordUserID(char.DiscordUserID).DoAndReturn(func(_ string) (*repository.
		Character, error) {
		return nil, nil
	}).MaxTimes(1)

	s := commands.NewMockDiscordConnector(ctrl)

	message := discordgo.Message{
		ID:        "428",
		ChannelID: "43",
		GuildID:   "",
		Content:   "Test content",
	}

	messageCreate := discordgo.MessageCreate{Message: &message}

	generator := create.CharCommandGenerator{Repo: m}

	charCreateCommand, err := generator.CreateCommand(s, &messageCreate, "TestChar", "wrong", "1234")
	require.NoError(t, err, "should create char command")
	assert.Equal(t, helpers.WrongClassGiven, charCreateCommand.Payload.Answer)
}
