package create

type Receiver struct{}

func (r *Receiver) Answer(p *CreateCommandPayload) {
	p.Session().ChannelMessageSend(p.Message.ChannelID, p.Answer)
}
