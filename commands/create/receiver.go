package create

type Receiver struct{}

func (r *Receiver) Answer(p *CharacterCreateCommandPayload) {
	p.Session().ChannelMessageSend(p.Message.ChannelID, p.Answer)
}
