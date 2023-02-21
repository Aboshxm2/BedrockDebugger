package command

import (
	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/bwmarrin/discordgo"
)

type Jump struct {
}

func (Jump) Name() string {
	return "jump"
}

func (Jump) Description() string {
	return "jump once"
}

func (Jump) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	se, ok := session.GetSession(m.Author.ID)

	if !ok {
		s.ChannelMessageSendReply(m.ChannelID, "**You don't have an open session**", m.Reference())
		return
	}

	se.Jump()

	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
}
