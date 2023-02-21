package command

import (
	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/bwmarrin/discordgo"
)

type Move struct {
}

func (Move) Name() string {
	return "move"
}

func (Move) Description() string {
	return "move one block to specific direction"
}

func (Move) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	se, ok := session.GetSession(m.Author.ID)

	if !ok {
		s.ChannelMessageSendReply(m.ChannelID, "**You don't have an open session**", m.Reference())
		return
	}

	if len(args) < 1 {
		s.ChannelMessageSendReply(m.ChannelID, "**Usage: move <direction>**", m.Reference())
		return
	}

	direction := args[0]
	if direction != "east" && direction != "west" && direction != "north" &&
		direction != "south" && direction != "up" && direction != "down" {
		s.ChannelMessageSendReply(m.ChannelID, "**direction must be on of these values: <east | west | north | south | west | up | down>**", m.Reference())
		return
	}

	se.Move(direction)

	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
}
