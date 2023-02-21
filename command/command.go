package command

import "github.com/bwmarrin/discordgo"

type Command interface {
	Name() string
	Description() string
	Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string)
}
