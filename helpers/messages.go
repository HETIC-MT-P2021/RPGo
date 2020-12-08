package helpers

import (
	"github.com/HETIC-MT-P2021/RPGo/commands"
	"github.com/bwmarrin/discordgo"
)

// Messages sent to discord API
const (
	CharAlreadyExists       = "You already have a character!"
	CharSuccessfullyCreated = "%s successfully created!"
	WrongClassGiven         = "You have to choose an existing class among : ['wizard', 'rogue', 'knight']!"
	GenericUserError        = "An issue occurred, please try again later."
	CharacterPresentation   = "Hello dear friend, my name is %s and I am a proud %v"
	CharacterDoesNotExist   = "Character does not exist, create your character first"
)

var errorEmbed = &discordgo.MessageEmbed{
	Title:       "Oops!",
	Description: "An issue occurred, please try again later",
	Color:       ErrorRedColor,
}

//SendGenericErrorMessage sends an error message to end user
func SendGenericErrorMessage(session commands.DiscordConnector, channelID string) {
	session.ChannelMessageSend(channelID, GenericUserError)
}

//SendGenericErrorEmbedMessage sends an error message to end user with embed style
func SendGenericErrorEmbedMessage(session commands.DiscordConnector, channelID string) {
	session.ChannelMessageSendEmbed(channelID, errorEmbed)
}
