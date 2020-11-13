package helpers

import "github.com/HETIC-MT-P2021/RPGo/commands"

// Messages sent to discord API
const (
	CharAlreadyExists       = "You already have a character!"
	CharSuccessfullyCreated = "%s successfully created!"
	GenericUserError        = "An issue occured, please try again later."
)

//SendGenericErrorMessage sends an error message to end user
func SendGenericErrorMessage(session commands.DiscordConnector, channelID string) {
	session.ChannelMessageSend(channelID, GenericUserError)
}
