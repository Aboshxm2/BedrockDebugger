package command

import (
	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/bwmarrin/discordgo"
)

type Sneak struct {
}

func (Sneak) Name() string {
	return "sneak"
}

func (Sneak) Description() string {
	return "toggle sneak"
}

func (Sneak) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	se, ok := session.GetSession(m.Author.ID)

	if !ok {
		s.ChannelMessageSendReply(m.ChannelID, "**You don't have an open session**", m.Reference())
		return
	}

	se.Sneak()

	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
}
