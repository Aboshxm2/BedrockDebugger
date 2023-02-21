package session

import (
	"github.com/Aboshxm2/BedrockDebugger/bedrock"
	"github.com/bwmarrin/discordgo"
)

type Session struct {
	c  *bedrock.Client
	s  *discordgo.Session
	ch *discordgo.Channel

	sId string
}

func NewSession(s *discordgo.Session, ch *discordgo.Channel, sId string) *Session {
	return &Session{s: s, ch: ch, sId: sId}
}

func (s *Session) InitClient(c *bedrock.Client) {
	s.c = c

	s.c.OnSpawn = func() {
		s.s.ChannelMessageSend(s.ch.ID, "The player has spawned")
	}

	s.c.OnMessage = func(message string) {
		s.s.ChannelMessageSend(s.ch.ID, message)
	}

	s.c.OnClose = func(reason string) {
		s.c.RainbowSkin = false

		s.s.ChannelMessageSend(s.ch.ID, "The connection has closed, Reason: "+reason)
		s.closeThread()
		RemoveSession(s.sId)
	}
}

func (s *Session) Chat(message string) {
	s.c.Chat(message)
}

func (s *Session) Disconnect(reason string) {
	s.c.Disconnect(reason)
}

func (s *Session) closeThread() {
	t := true
	s.s.ChannelEditComplex(s.ch.ID, &discordgo.ChannelEdit{
		Archived: &t, // yes as you can see, this is a pointer to boolean type
		Locked:   &t,
	})
}

func (s *Session) Attack(player string) {
	s.c.Attack(player)
}

func (s *Session) Move(direction string) {
	s.c.Move(direction)
}

func (s *Session) Jump() {
	s.c.Jump()
}

func (s *Session) Sneak() {
	s.c.Sneak()
}

func (s *Session) Rainbow() {
	s.c.Rainbow()
}
