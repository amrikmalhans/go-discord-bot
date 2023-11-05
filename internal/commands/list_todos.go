package commands

import (
	"context"
	"fmt"

	"github.com/amrikmalhans/go-discord-bot.git/internal/db"
	"github.com/bwmarrin/discordgo"
)

type Todo struct {
	ID        int
	UserId    string
	Item      string
	CreatedAt string
	UpdateAt  string
}

func ListTodos(s *discordgo.Session, i *discordgo.InteractionCreate) {

	data := i.ApplicationCommandData()
	userId := i.Member.User.ID

	db := db.InitDB()

	switch data.Name {
	case "list-todos":

		rows, err := db.Query(context.Background(), "SELECT item FROM todos WHERE user_id = $1", userId)
		if err != nil {
			panic(err)
		}

		// Create an empty slice of MessageEmbedField pointers
		fields := []*discordgo.MessageEmbedField{}

		index := 0
		components := make([]discordgo.MessageComponent, 0)
		for rows.Next() {
			index++
			var todoItem Todo
			err := rows.Scan(&todoItem.Item)
			if err != nil {
				panic(err)
			}
			fields = append(fields, &discordgo.MessageEmbedField{
				Name:  fmt.Sprintf("Todo #%d", index),
				Value: todoItem.Item,
			})

			button := discordgo.Button{
				Label: "Done",
				Style: discordgo.SuccessButton,
				CustomID: fmt.Sprintf("todo-%d", todoItem.ID),
			}

			components = append(components, button)
		}

		

		embed := &discordgo.MessageEmbed{
			Title:       "Your todo items",
			Description: "Here are your todo items",
			Color:       0x00ff00,
			Fields:      fields,
		}

		err = s.InteractionRespond(
			i.Interaction,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Embeds: []*discordgo.MessageEmbed{embed},
					Components: components,
				},
			},
		)
		if err != nil {
			panic(err)
		}
	}
}
