package disc

import "github.com/disgoorg/disgo/discord"

type CommandName string

const (
	InfoCmdName CommandName = "info"
	EchoCmdName CommandName = "echo"
)

var (
	DiscordCmds = map[CommandName]discord.ApplicationCommandCreate{
		InfoCmdName: discord.SlashCommandCreate{
			Name:        "info",
			Description: "display info about the reap bot",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionBool{
					Name:        "exhaustive",
					Description: "Overflow you with probably useless info? :P",
				},
				discord.ApplicationCommandOptionString{
					Name:        "free",
					Description: "Put whatever message you want here",
				},
			},
		},
		EchoCmdName: discord.SlashCommandCreate{
			Name:        "echo",
			Description: "test the reap bot is working!",
			Options: []discord.ApplicationCommandOption{
				discord.ApplicationCommandOptionString{
					Name:        "message",
					Description: "message to echo back",
				},
			},
		},
	}
)
