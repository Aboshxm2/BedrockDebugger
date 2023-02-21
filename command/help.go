package command

import "github.com/bwmarrin/discordgo"

type Help struct {
	Commands []Command
}

func (Help) Name() string {
	return "help"
}

func (Help) Description() string {
	return "" // this line will never executed
}

func (h Help) Execute(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	msg := "```\n"

	for _, command := range h.Commands {
		msg += command.Name() + " - " + command.Description() + "\n"
	}

	msg +=
		"```\n" +
			"**Github:** <https://github.com/Aboshxm2/BedrockDebugger>\n" +
			"**Support:** <https://discord.gg/SRHJAKFfzc>"

	s.ChannelMessageSend(m.ChannelID, msg)
}
