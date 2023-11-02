package commands

import (
	"context"
	"fmt"

	"github.com/amrikmalhans/go-discord-bot.git/internal/db"
	"github.com/bwmarrin/discordgo"
)

func AddTodo(s *discordgo.Session, i *discordgo.InteractionCreate) {

	data := i.ApplicationCommandData()
	userId := i.Member.User.ID

	db := db.InitDB()

	switch data.Name {
	case "add-todo":

		todoItem := data.Options[0].StringValue()

		_, err := db.Exec(context.Background(), "INSERT INTO todos (user_id, item) VALUES ($1, $2)", userId, todoItem)
		if err != nil {
			panic(err)
		}

		err = s.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: fmt.Sprintf("Added todo item: %s", todoItem),
				},
			},
		)
		if err != nil {
			panic(err)
		}
	}
}
