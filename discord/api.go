package discord

import (
	"time"

	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v5"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
)

var c *disgord.Client

func apiInit(client *disgord.Client) {
	c = client
}

func API(link string) {
	s, err := snowflake.GetSnowflake(config.API_Channel)
	if common.CheckErr(err, "get snowflake") {
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
