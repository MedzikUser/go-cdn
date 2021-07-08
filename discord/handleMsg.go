package discord

import (
	"time"

	"github.com/andersfylling/disgord"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
	"gitlab.com/gaming0skar123/go/cdn/database"
	"gitlab.com/gaming0skar123/go/cdn/imgur"
)

func handleMsg(s disgord.Session, data *disgord.MessageCreate) {
	m := data.Message

	findVaue := &database.Channel{
		ID:      m.ChannelID.String(),
		GuildID: m.GuildID.String(),
	}

	if !database.FindChannel(findVaue) {
		return
	}

	err := s.Channel(m.ChannelID).TriggerTypingIndicator()
	common.CheckErr(err, "TriggerTypingIndicator()")

	avatarUrl, _ := m.Author.AvatarURL(64, true)

	var url string

	if len(m.Attachments) > 0 {
		url = m.Attachments[0].URL
	} else {
		url = m.Content
	}

	i, err := imgur.UploadFromURL(url)
	if common.CheckErr(err, "UploadFromURL("+url+")") {
		mBot, err := m.Reply(ctx, s, disgord.Embed{
			Title:       "Error",
			Description: err.Error(),
			Timestamp:   m.Timestamp,
			Color:       config.Embed_Color,
			Footer: &disgord.EmbedFooter{
				Text:    m.Author.Tag(),
				IconURL: avatarUrl,
			},
		})
		if !common.CheckErr(err, "send message") {
			deleteMsg(s, mBot, 5*time.Second)
		}

		deleteMsg(s, m, 3*time.Second)

		return
	}

	_, err = m.Reply(ctx, s, disgord.Embed{
		Title:     i.Link,
		Timestamp: m.Timestamp,
		Color:     config.Embed_Color,
		Footer: &disgord.EmbedFooter{
			Text:    m.Author.Tag(),
			IconURL: avatarUrl,
		},
		Image: &disgord.EmbedImage{
			URL: i.Link,
		},
	})
	common.CheckErr(err, "send message")

	deleteMsg(s, m, 2*time.Second)
}
