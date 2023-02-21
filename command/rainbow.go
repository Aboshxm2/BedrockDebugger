package command

import (
	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/bwmarrin/discordgo"
)

type Rainbow struct {
}

func (Rainbow) Name() string {
	return "rainbow"
}

func (Rainbow) Description() string {
	return "change skin in a strange way"
}

func (Rainbow) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	se, ok := session.GetSession(m.Author.ID)

	if !ok {
		s.ChannelMessageSendReply(m.ChannelID, "**You don't have an open session**", m.Reference())
		return
	}

	se.Rainbow()

	s.MessageReactionAdd(m.ChannelID, m.ID, "🌈")
}
