package ping

import (
	"github.com/bwmarrin/discordgo"
)

type PingCommand struct {
	Receiver *Receiver
	Payload  *PingCommandPayload
}

type PingCommandPayload struct {
	Answer  string
	Session *discordgo.Session
	Message *discordgo.MessageCreate
}

func MakePingCommand(s *discordgo.Session, m *discordgo.MessageCreate) *PingCommand {

	return &PingCommand{
		Receiver: &Receiver{},
		Payload: &PingCommandPayload{
			Answer:  "Pong!",
			Session: s,
			Message: m,
		},
	}
}

func (c *PingCommand) Execute() {
	c.Receiver.Answer(c.Payload)
}
