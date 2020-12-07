package help

import (
	"github.com/HETIC-MT-P2021/RPGo/commands"
	customenv "github.com/HETIC-MT-P2021/RPGo/env"
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
	session commands.DiscordConnectorMessageEmbed
	Message *discordgo.MessageCreate
}

//Command a help command
func MakeCommand(c commands.DiscordConnectorMessageEmbed, m *discordgo.MessageCreate) *Command {
	field := &discordgo.MessageEmbedField{
		Name:   "New player ?",
		Value:  "Try to type `" + customenv.DiscordPrefix + "create {characterName} {characterClass}` to create your own character !",
		Inline: true,
	}

	var fields = make([]*discordgo.MessageEmbedField, 0)

	fields = append(fields, field)
	lightGreenDecimal := 6744188

	answer := &discordgo.MessageEmbed{
		Title:       "RPGo help",
		Description: "Welcome to RPGo, here's all the help you need to play with the bot !",
		Color:       lightGreenDecimal,
		Footer:      nil,
		Image:       nil,
		Thumbnail:   nil,
		Video:       nil,
		Provider:    nil,
		Author:      nil,
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
func (p *CommandPayload) Session() commands.DiscordConnectorMessageEmbed {
	return p.session
}
