package disc

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
)

func handleMessageCreate(log *slog.Logger) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.MessageCreate) {
		log.Info("MESSAGE_CREATE")
	})
}

func handleDmMessageCreate(log *slog.Logger) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.DMMessageCreate) {
		log.Info("DM_MESSAGE_CREATE")
	})
}

func handleDMMessageUpdate(log *slog.Logger) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.DMMessageUpdate) {
		log.Info("DM_MESSAGE_UPDATE")
	})
}

func handleDMMessageDelete(log *slog.Logger) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.DMMessageDelete) {
		log.Info("DM_MESSAGE_DELETE")
	})
}

func handleApplicationCommandInteractionCreate(log *slog.Logger) bot.EventListener {
	return bot.NewListenerFunc(func(e *events.ApplicationCommandInteractionCreate) {
		log.Info("APPLICATION_COMMAND_INTERACTION_CREATE")
	})
}
