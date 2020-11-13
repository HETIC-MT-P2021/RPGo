package commands

import "github.com/bwmarrin/discordgo"

//DiscordConnector interface for sessions
type DiscordConnector interface {
	Create(s *discordgo.Session, m *discordgo.MessageCreate, name string, userID string)
}
