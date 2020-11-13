package main

import (
	"flag"
	"fmt"
	"github.com/HETIC-MT-P2021/RPGo/commands/create"
	"github.com/HETIC-MT-P2021/RPGo/commands/ping"
	"github.com/HETIC-MT-P2021/RPGo/database"
	"github.com/HETIC-MT-P2021/RPGo/repository"
	"github.com/bwmarrin/discordgo"
	"github.com/caarlos0/env/v6"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// Variables used for command line parameters
var (
	Token string
)

var discordPrefix = "&"

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

//DiscordConfig dto
type DiscordConfig struct {
	Token string `env:"TOKEN"`
}

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	log.Printf("connected to database")

	dsConfig := DiscordConfig{}
	if err := env.Parse(&dsConfig); err != nil {
		return
	}

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + dsConfig.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the message is "ping" reply with "Pong!"
	if m.Content == (discordPrefix + "ping") {
		pingCommand := ping.MakePingCommand(s, m)
		pingCommand.Execute()
	}

	if strings.HasPrefix(m.Content, discordPrefix+"create") {
		if len(strings.Fields(m.Content)) > 1 {
			args := make([]string, 0)
			for _, word := range strings.Fields(m.Content) {
				args = append(args, word)
			}

			commandGenerator := create.CharCommandGenerator{
				Repo: &repository.CharacterRepository{
					Conn: database.DBCon,
				}}

			createCommand := commandGenerator.Create(s, m, args[1], m.Author.ID)
			createCommand.Execute()

			return
		}
		s.ChannelMessageSend(m.ChannelID, "No name given! Try `&create {characterName}`")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == (discordPrefix + "pong") {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
