package discord

import (
	"time"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v5"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
	"gitlab.com/gaming0skar123/go/cdn/database"
)

var c *disgord.Client

func apiInit(client *disgord.Client) {
	c = client
}

func API(link string) {
	channel, err := database.FindAPIChannel()
	if common.CheckErr(err, "FindAPIChannel()") {
		return
	}

	s, err := snowflake.GetSnowflake(channel.ID)
	if common.CheckErr(err, "GetSnowflake("+channel.ID+")") {
		return
	}

	_, err = c.SendMsg(s, disgord.Embed{
		Title: link,
		Timestamp: disgord.Time{
			Time: time.Now().UTC(),
		},
		Color: config.Embed_Color,
		Image: &disgord.EmbedImage{
			URL: link,
		},
	})

	if common.CheckErr(err, "send msg") {
		return
	}
}
