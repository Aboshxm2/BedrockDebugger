package command

import (
	"strings"

	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/bwmarrin/discordgo"
)

type Chat struct {
}

func (Chat) Name() string {
	return "chat"
}

func (Chat) Description() string {
	return "send chat message"
}

func (Chat) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	se, ok := session.GetSession(m.Author.ID)

	if !ok {
		s.ChannelMessageSendReply(m.ChannelID, "**You don't have an open session**", m.Reference())
		return
	}

	if len(args) < 1 {
		s.ChannelMessageSendReply(m.ChannelID, "**Usage: say <message>**", m.Reference())
		return
	}

	se.Chat(strings.Join(args, " "))

	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
}
