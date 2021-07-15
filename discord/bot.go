package discord

import (
	"context"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
)

var ctx = context.Background()

var log = common.Log

func Bot() {
	client := disgord.New(disgord.Config{
		ProjectName: "CDN",
		BotToken:    config.Bot_Token,
		Logger:      log,
		RejectEvents: []string{
			disgord.EvtTypingStart,

			disgord.EvtPresenceUpdate,
			disgord.EvtGuildMemberAdd,
			disgord.EvtGuildMemberUpdate,
			disgord.EvtGuildMemberRemove,
		},
		// DMIntents: disgord.IntentDirectMessages | disgord.IntentDirectMessageReactions | disgord.IntentDirectMessageTyping,
	})

	defer func() {
		err := client.Gateway().StayConnectedUntilInterrupted()
		common.CheckErr(err, "stay bot connection")
	}()

	filter, _ := std.NewMsgFilter(ctx, client)

	client.Gateway().WithMiddleware(
		filter.NotByBot, // ignore bot messages
	).MessageCreate(handleMsg)

	apiInit(client)

	client.Gateway().BotReady(func() {
		log.Info("Bot is ready!")
	})
}
