package discord

import (
	"time"

	"github.com/andersfylling/disgord"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
	"gitlab.com/gaming0skar123/go/cdn/imgur"
)

func uploadImg(s disgord.Session, m *disgord.Message, url string) {
	avatarUrl, _ := m.Author.AvatarURL(64, true)

	i, err := imgur.UploadFromURL(url)
	if common.CheckErr(err, "upload img") {
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
		if !common.CheckErr(err, "send msg") {
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
	common.CheckErr(err, "send msg")
}
