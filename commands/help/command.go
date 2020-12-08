package help

import (
	"github.com/HETIC-MT-P2021/RPGo/commands"
	customenv "github.com/HETIC-MT-P2021/RPGo/env"
	"github.com/HETIC-MT-P2021/RPGo/helpers"
	"github.com/bwmarrin/discordgo"
)

//Command model for help command
type Command struct {
	Receiver *Receiver
	payload  *CommandPayload
}

//CommandPayload model for information on help command
type CommandPayload struct {
	Answer  *discordgo.MessageEmbed
	session commands.DiscordConnector
	Message *discordgo.MessageCreate
}

//MakeCommand a help command
func MakeCommand(c commands.DiscordConnector, m *discordgo.MessageCreate) *Command {
	field := &discordgo.MessageEmbedField{
		Name:   "New player ?",
		Value:  "Try to type `" + customenv.DiscordPrefix + "create {characterName} {characterClass}` to create your own character !",
		Inline: true,
	}

	var fields = make([]*discordgo.MessageEmbedField, 0)

	// Here possibility to add help fields as the project grows
	fields = append(fields, field)

	answer := &discordgo.MessageEmbed{
		Title:       "RPGo help",
		Description: "Welcome to RPGo, here's all the help you need to play with the bot !",
		Color:       helpers.LightGreenDecimal,
		Fields:      fields,
	}

	return &Command{
		Receiver: &Receiver{},
		payload: &CommandPayload{
			Answer:  answer,
			session: c,
			Message: m,
		},
	}
}

//Execute command with all its information
func (c *Command) Execute() {
	c.Receiver.Answer(c.Payload())
}

//Payload returns CreateCommand payload
func (c *Command) Payload() *CommandPayload {
	return c.payload
}

//Session returns CreateCommand session
func (p *CommandPayload) Session() commands.DiscordConnector {
	return p.session
}
