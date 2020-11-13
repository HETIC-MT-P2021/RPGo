package commands

import "github.com/bwmarrin/discordgo"

//Command interface (command pattern)
type Command interface {
	Execute() error
	Payload() *CommandPayload
}

//CommandPayload pattern for command information
type CommandPayload interface {
	Session() *discordgo.Session
}
