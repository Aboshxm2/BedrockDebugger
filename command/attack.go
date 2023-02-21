package command

import (
	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/bwmarrin/discordgo"
)

type Attack struct {
}

func (Attack) Name() string {
	return "attack"
}

func (Attack) Description() string {
	return "hit player once"
}

func (Attack) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	se, ok := session.GetSession(m.Author.ID)

	if !ok {
		s.ChannelMessageSendReply(m.ChannelID, "**You don't have an open session**", m.Reference())
		return
	}

	if len(args) < 1 {
		s.ChannelMessageSendReply(m.ChannelID, "**Usage: attack <player>**", m.Reference())
		return
	}

	se.Attack(args[0])

	s.MessageReactionAdd(m.ChannelID, m.ID, "✅")
}
