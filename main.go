package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/Aboshxm2/BedrockDebugger/auth"
	"github.com/Aboshxm2/BedrockDebugger/command"
	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/Aboshxm2/BedrockDebugger/translate"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var PREFIX string

var commands = []command.Command{
	command.Ping{},
	command.Query{},
	command.Join{},
	command.Leave{},
	command.Chat{},
	command.Attack{},
	command.Move{},
	command.Sneak{},
	command.Jump{},
	command.Rainbow{},
}

func main() {
	// Initialization xbox token.
	auth.TokenSource()

	// Initialization language used to translate (TextTypeTranslation)[https://github.com/Sandertv/gophertunnel/blob/master/minecraft/protocol/packet/text.go#L10].
	// You can get this file from https://github.com/Mojang/bedrock-samples/blob/main/resource_pack/texts/en_US.lang.
	translate.LoadLang("en_US.lang")

	commands = append(commands, command.Help{
		Commands: commands,
	})

	godotenv.Load(".env")

	PREFIX = os.Getenv("PREFIX")

	discord, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}

	discord.Identify.Intents = discordgo.IntentsGuildMessages

	discord.AddHandler(messageCreate)

	discord.Identify.Presence = discordgo.GatewayStatusUpdate{
		Game: discordgo.Activity{
			Name: PREFIX + "help",
			Type: discordgo.ActivityTypeGame,
		},
	}

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()

	for _, session := range session.GetSessions() {
		session.Disconnect("The bot has shutdown")
	}
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, PREFIX) {
		return
	}

	args := strings.Split(strings.TrimPrefix(m.Content, PREFIX), " ")
	if len(args) < 1 {
		return
	}

	commandName := args[0]
	args = args[1:]

	for _, command := range commands {
		if command.Name() == commandName {
			command.Execute(s, m, args)
			break
		}
	}
}
