package commands

import "github.com/bwmarrin/discordgo"

type Command interface {
	Execute() error
	Payload() *CommandPayload
}

type CommandPayload interface {
	Session() *discordgo.Session
}
