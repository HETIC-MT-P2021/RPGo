package create

//Receiver to receive command and send according message
type Receiver struct{}

//Answer sends channel message to discord bot according to the command
func (r *Receiver) Answer(p *CharacterCreateCommandPayload) {
	p.Session().ChannelMessageSend(p.Message.ChannelID, p.Answer)
}
