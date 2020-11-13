package ping

import (
	"github.com/bwmarrin/discordgo"
)

//Command model for Ping Command
type Command struct {
	Receiver *Receiver
	payload  *CommandPayload
}

//CommandPayload information on Ping Command
type CommandPayload struct {
	Answer  string
	session *discordgo.Session
	Message *discordgo.MessageCreate
}

//MakePingCommand makes a ping command
func MakePingCommand(s *discordgo.Session, m *discordgo.MessageCreate) *Command {

	return &Command{
		Receiver: &Receiver{},
		payload: &CommandPayload{
			Answer:  "Pong!",
			session: s,
			Message: m,
		},
	}
}

//Execute command with all its information
func (c *Command) Execute() {
	c.Receiver.Answer(c.Payload())
}

//Payload returns PingCommand payload
func (c *Command) Payload() *CommandPayload {
	return c.payload
}

//Session returns PingCommandPayload discord session
func (p *CommandPayload) Session() *discordgo.Session {
	return p.session
}
