package ping

import (
	"github.com/bwmarrin/discordgo"
)

type PingCommand struct {
	Receiver *Receiver
	payload  *PingCommandPayload
}

type PingCommandPayload struct {
	Answer  string
	session *discordgo.Session
	Message *discordgo.MessageCreate
}

func MakePingCommand(s *discordgo.Session, m *discordgo.MessageCreate) *PingCommand {

	return &PingCommand{
		Receiver: &Receiver{},
		payload: &PingCommandPayload{
			Answer:  "Pong!",
			session: s,
			Message: m,
		},
	}
}

func (c *PingCommand) Execute() {
	c.Receiver.Answer(c.Payload())
}

func (c *PingCommand) Payload() *PingCommandPayload {
	return c.payload
}

func (p *PingCommandPayload) Session() *discordgo.Session {
	return p.session
}
