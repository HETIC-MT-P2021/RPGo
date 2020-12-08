package help

import (
	"github.com/HETIC-MT-P2021/RPGo/commands"
	customenv "github.com/HETIC-MT-P2021/RPGo/env"
	"github.com/bwmarrin/discordgo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var field = &discordgo.MessageEmbedField{
	Name:   "New player ?",
	Value:  "Try to type `" + customenv.DiscordPrefix + "create {characterName} {characterClass}` to create your own character !",
	Inline: true,
}

var fields = make([]*discordgo.MessageEmbedField, 0)

func TestHelpCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fields = append(fields, field)

	expectedHelpAnswer := &discordgo.MessageEmbed{
		Title:       "RPGo help",
		Description: "Welcome to RPGo, here's all the help you need to play with the bot !",
		Color:       6744188,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
		Fields:      fields,
	}

	s := commands.NewMockDiscordConnector(ctrl)

	message := discordgo.Message{
		ID:        "428",
		ChannelID: "43",
		GuildID:   "",
		Content:   "Test content",
	}

	messageCreate := discordgo.MessageCreate{Message: &message}

	helpCommand := MakeCommand(s, &messageCreate)
	assert.Equal(t, expectedHelpAnswer, helpCommand.payload.Answer)
}
