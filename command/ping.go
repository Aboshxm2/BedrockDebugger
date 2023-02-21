package command

import (
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Ping struct {
}

func (Ping) Name() string {
	return "ping"
}

func (Ping) Description() string {
	return "get bot ping"
}

func (Ping) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	start := time.Now()

	msg, err := s.ChannelMessageSend(m.ChannelID, "Pong!")
	if err != nil {
		return
	}

	heartbeat := strconv.FormatInt(s.HeartbeatLatency().Milliseconds(), 10)
	latency := strconv.FormatInt(time.Since(start).Milliseconds(), 10)

	s.ChannelMessageEdit(msg.ChannelID, msg.ID, "Pong!\nHeartbeat="+heartbeat+"ms, Latency="+latency+"ms")
}
