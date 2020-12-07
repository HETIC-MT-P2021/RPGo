package helpers

import "github.com/HETIC-MT-P2021/RPGo/commands"

// Messages sent to discord API
const (
	CharAlreadyExists       = "You already have a character!"
	CharSuccessfullyCreated = "%s successfully created!"
	WrongClassGiven         = "You have to choose an existing class!"
	GenericUserError        = "An issue occurred, please try again later."
	CharacterPresentation   = "Hello dear friend, my name is %s and I am a proud %s"
)

//SendGenericErrorMessage sends an error message to end user
func SendGenericErrorMessage(session commands.DiscordConnector, channelID string) {
	session.ChannelMessageSend(channelID, GenericUserError)
}
