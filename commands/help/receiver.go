package help

//Receiver to receive command and send according message
type Receiver struct{}

//Answer sends channel message to discord bot according to the command
func (r *Receiver) Answer(p *CommandPayload) {
	p.Session().ChannelMessageSendEmbed(p.Message.ChannelID, p.Answer)
}
