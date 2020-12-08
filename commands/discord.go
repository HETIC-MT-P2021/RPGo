package commands

import "github.com/bwmarrin/discordgo"

//DiscordConnectorMessage interface for sessions
type DiscordConnectorMessage interface {
	ChannelMessageSend(string, string) (*discordgo.Message, error)
}

//DiscordConnectorMessageEmbed interface for sessions
type DiscordConnectorMessageEmbed interface {
	ChannelMessageSendEmbed(string, *discordgo.MessageEmbed) (*discordgo.Message, error)
}
