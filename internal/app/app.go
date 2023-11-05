package app

import (
	"log"
	"os"
	"os/signal"

	"github.com/amrikmalhans/go-discord-bot.git/internal/commands"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

type App struct {
	Bot *discordgo.Session
}

func New() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		log.Fatal("No bot token provided")
	}

	guildID := os.Getenv("GUILD_ID")
	applicationID := os.Getenv("APPLICATION_ID")

	s, _ := discordgo.New("Bot " + botToken)

	_, err = s.ApplicationCommandBulkOverwrite(applicationID, guildID, []*discordgo.ApplicationCommand{
		{
			Name:        "add-todo",
			Description: "Add a to-do item",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "todo",
					Description: "The todo item you want to create",
					Required:    true,
				},
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:		"due-date",
					Description: "The due date for the todo item",
					Required:    false,
				},
			},
		},
		{
			Name:        "list-todos",
			Description: "List all todo items",
		},
	})

	s.Identify.Intents =  discordgo.IntentsGuildMessages

	if err != nil {
		log.Fatalf("Error setting up commands: %v", err)
	}

	commands.Register(s)

	return &App{
		Bot: s,
	}
}

func (a *App) Run() error {

	err := a.Bot.Open()
	if err != nil {
		return err
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = a.Bot.Close()
	if err != nil {
		return err
	}

	return nil
}
