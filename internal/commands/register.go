package commands

import (
	"github.com/bwmarrin/discordgo"
)

func Register(s *discordgo.Session) {
	s.AddHandler(AddTodo)
	s.AddHandler(ListTodos)
	s.AddHandler(Message)
}
