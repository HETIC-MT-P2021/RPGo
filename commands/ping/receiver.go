package ping

type Receiver struct{}

func (r *Receiver) Answer(p *PingCommandPayload) {
	p.Session().ChannelMessageSend(p.Message.ChannelID, p.Answer)
}
