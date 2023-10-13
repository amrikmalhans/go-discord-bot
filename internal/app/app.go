package app

import (
	"log"
	"os"
	"os/signal"

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
			Name:        "hello-world",
			Description: "Showcase of a basic slash command",
		},
	})

	if err != nil {
		log.Fatalf("Error setting up commands: %v", err)
	}

	s.AddHandler(func(
		s *discordgo.Session,
		i *discordgo.InteractionCreate,
	) {
		data := i.ApplicationCommandData()
		switch data.Name {
		case "hello-world":
			err := s.InteractionRespond(
				i.Interaction,
				&discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: "Hello world!",
					},
				},
			)
			if err != nil {
				panic(err)
			}
		}
	})

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
