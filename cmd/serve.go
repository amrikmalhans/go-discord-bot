package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

// Bot parameters
var (
	GuildID       = "848498442744627241"
	BotToken      = "MTE1Njc4MTI4ODIzNDE1NjA0Mg.GXqX9D.pGxeCAVytC2I875LghxJqDn4Y-fud1zP_qZVBQ"
	ApplicationID = "1156781288234156042"
)

func main() {

	s, _ := discordgo.New("Bot " + BotToken)

	_, err := s.ApplicationCommandBulkOverwrite(ApplicationID, GuildID, []*discordgo.ApplicationCommand{
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
	err = s.Open()
	if err != nil {
		panic(err)
	}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	log.Println("Press Ctrl+C to exit")
	<-stop

	err = s.Close()
	if err != nil {
		panic(err)
	}

}
