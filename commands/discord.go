package commands

import "github.com/bwmarrin/discordgo"

//DiscordConnector interface for sessions
type DiscordConnector interface {
	ChannelMessageSend(string, string) (*discordgo.Message, error)
	ChannelMessageSendEmbed(string, *discordgo.MessageEmbed) (*discordgo.Message, error)
}
