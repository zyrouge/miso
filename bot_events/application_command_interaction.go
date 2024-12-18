package bot_events

import (
	"log/slog"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"me.zyrouge.tofu/core"
)

func NewTofuApplicationCommandInteractionEvent() core.TofuEvent {
	return core.NewTofuEvent(func(tofu *core.Tofu, event *events.ApplicationCommandInteractionCreate) {
		if !tofu.FilteredGuilds.IsWhitelisted(event.GuildID()) {
			return
		}
		data := event.SlashCommandInteractionData()
		if command, ok := tofu.Commands[data.CommandName()]; ok {
			response := command.Invoke(tofu, event)
			if err := event.Respond(discord.InteractionResponseTypeCreateMessage, response); err != nil {
				slog.Error("application command interaction event respond failed: " + err.Error())
			}
		}
	})
}