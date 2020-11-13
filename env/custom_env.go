package env

//DiscordPrefix defines what the user uses has a prefix on Discord
var DiscordPrefix = "&"

//DiscordConfig dto
type DiscordConfig struct {
	Token string `env:"TOKEN"`
}
