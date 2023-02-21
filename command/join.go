package command

import (
	"strings"

	"github.com/Aboshxm2/BedrockDebugger/bedrock"
	"github.com/Aboshxm2/BedrockDebugger/session"
	"github.com/bwmarrin/discordgo"
)

type Join struct {
}

func (Join) Name() string {
	return "join"
}

func (Join) Description() string {
	return "join minecraft server"
}

func (Join) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if _, ok := session.GetSession(m.Author.ID); ok {
		s.ChannelMessageSendReply(m.ChannelID, "**You can't open more than one session at the same time.**", m.Reference())
		return
	}

	if len(args) < 1 || len(args) > 2 {
		s.ChannelMessageSendReply(m.ChannelID, "**Usage: join <address>**", m.Reference())
		return
	}

	address := strings.Join(args, ":")
	if len(args) == 1 && !strings.Contains(address, ":") {
		address += ":19132"
	}

	msg, err := s.ChannelMessageSendReply(m.ChannelID, "**Connecting to "+address+"**", m.Reference())
	if err != nil {
		return
	}

	c, err := bedrock.NewClient(address)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, "**Error: "+err.Error()+"**", m.Reference())
		return
	}

	thread, err := s.MessageThreadStart(m.ChannelID, msg.ID, address, 60)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, "**Error: "+err.Error()+"**", m.Reference())
		return
	}

	se := session.NewSession(s, thread, m.Author.ID)

	session.AddSession(m.Author.ID, se)

	se.InitClient(c)

	c.Connect()
}
