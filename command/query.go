package command

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/sandertv/gophertunnel/query"
)

type Query struct {
}

func (Query) Name() string {
	return "query"
}

func (Query) Description() string {
	return "get server details"
}

func (Query) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if len(args) < 1 || len(args) > 2 {
		s.ChannelMessageSendReply(m.ChannelID, "**Usage: query <address>**", m.Reference())
		return
	}

	address := strings.Join(args, ":")
	if len(args) == 1 && !strings.Contains(address, ":") {
		address += ":19132"
	}

	info, err := query.Do(address)
	if err != nil {
		s.ChannelMessageSendReply(m.ChannelID, "**Error: "+err.Error()+"**", m.Reference())
		return
	}

	msg := "```\n"

	msg += "host: " + info["hostname"] + "\n"
	msg += "ip: " + info["hostip"] + "\n"
	msg += "port: " + info["hostport"] + "\n"
	msg += "version: " + info["version"] + "\n"
	msg += "server_engine: " + info["server_engine"] + "\n"
	msg += "plugins: " + info["plugins"] + "\n"
	msg += "whitelist: " + info["whitelist"] + "\n"
	msg += "numplayers: " + info["numplayers"] + "\n"
	msg += "maxplayers: " + info["maxplayers"] + "\n"
	msg += "players: " + info["players"] + "\n"
	msg += "gametype: " + info["gametype"] + "\n"
	msg += "game_id: " + info["game_id"] + "\n"
	msg += "map: " + info["map"] + "\n"

	msg += "```"

	s.ChannelMessageSendReply(m.ChannelID, msg, m.Reference())
}
