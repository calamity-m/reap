package disc

import (
	"context"
	"fmt"
	"log/slog"
	"maps"
	"os"
	"os/signal"
	"slices"
	"syscall"

	"github.com/calamity-m/reap/harvest/internal/config"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
)

type DiscordBot struct {
	log    *slog.Logger
	client bot.Client
}

func (d *DiscordBot) Run() error {
	d.log.Debug("running")

	// Sync our commands as the first thing we do.
	if err := d.SyncGlobalCommands(context.TODO()); err != nil {
		return err
	}

	// Now register any listeners we have.
	d.client.AddEventListeners(handleMessageCreate(d.log))
	d.client.AddEventListeners(handleDmMessageCreate(d.log))
	d.client.AddEventListeners(handleDMMessageUpdate(d.log))
	d.client.AddEventListeners(handleDMMessageDelete(d.log))
	d.client.AddEventListeners(handleApplicationCommandInteractionCreate(d.log))

	if err := d.client.OpenGateway(context.TODO()); err != nil {
		return err
	}
	// Don't handle closing here, as our caller should be dealing with that.
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a Interrupt or Kill
	d.log.Info("harvet discord bot is running")
	<-sig

	return nil
}

func (d *DiscordBot) Close(ctx context.Context) {
	d.log.Debug("closing")
	d.client.Close(ctx)
}

// Retrieves a current list of registered commands from discord. Note that discord global command propogation
// to the discord client itself, takes a long time. (1 hour+)
func (d *DiscordBot) FetchGlobalCommands(ctx context.Context) ([]discord.ApplicationCommand, error) {
	cmds, err := d.client.Rest().GetGlobalCommands(d.client.ApplicationID(), true)

	if err != nil {
		return nil, err
	}

	return cmds, nil
}

// Syncs commands from DiscordCmds with discord as global commands. Any global commands that are registered, but not listed
// in the DiscordCmds var will be deleted.
func (d *DiscordBot) SyncGlobalCommands(ctx context.Context) error {
	// Relies on the doc from discord:
	// "Commands that do not already exist will count toward daily application command create limits."
	//
	// Due to this, we can just blatanly set our global commands at the start, and then just do simple
	// cleanup after.

	// First, create all of our commands
	set, err := d.client.Rest().SetGlobalCommands(
		d.client.ApplicationID(),
		// Magic fun of the stdlib iterators
		slices.Collect(maps.Values(DiscordCmds)),
	)
	if err != nil {
		return fmt.Errorf("failed to create global commands: %w", err)
	}

	d.log.InfoContext(ctx, "commands have been set", slog.Any("commands", set))

	// Now grab an updated list of commands so that we can delete unwanted ones
	commands, err := d.FetchGlobalCommands(ctx)
	if err != nil {
		return fmt.Errorf("failed to grab global commands: %w", err)
	}

	for _, command := range commands {
		_, ok := DiscordCmds[CommandName(command.Name())]
		if !ok {
			// We have a command that shouldn't exist
			d.log.Info("deleteting unwanted command", slog.Any("command", command))
			err := d.client.Rest().DeleteGlobalCommand(d.client.ApplicationID(), command.ID())
			if err != nil {
				// If we faild to delete, we'll just log and move onto the next, as we don't want to miss deleting the remaining
				// unwanted commands
				d.log.ErrorContext(ctx, "failed to delete unwanted command", slog.Any("command", command), slog.Any("err", err))
			}
		}
	}

	return nil
}

func NewDiscordBot(logger *slog.Logger, cfg *config.Config) (*DiscordBot, error) {
	if logger == nil {
		return nil, fmt.Errorf("nil logger not allowed")
	}

	if cfg == nil {
		return nil, fmt.Errorf("nil cfg not allowed")
	}

	client, err := disgo.New(
		cfg.Token,
		bot.WithGatewayConfigOpts(
			// set enabled intents
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildMessages,
				gateway.IntentDirectMessages,
				gateway.IntentGuildInvites,
				gateway.IntentMessageContent,
			),
		),
		bot.WithLogger(logger),
	)
	if err != nil {
		return nil, err
	}

	bot := &DiscordBot{log: logger, client: client}

	return bot, nil
}
