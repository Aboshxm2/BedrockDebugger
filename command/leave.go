package command

import (
	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/bwmarrin/discordgo"
)

type Leave struct {
}

func (Leave) Name() string {
	return "leave"
}

func (Leave) Description() string {
	return "leave from minecraft server"
}

func (Leave) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	se, ok := session.GetSession(m.Author.ID)

	if !ok {
		s.ChannelMessageSendReply(m.ChannelID, "**You don't have an open session**", m.Reference())
		return
	}

	se.Disconnect("user request")

	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
}
