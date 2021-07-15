package discord

import (
	"time"

	"github.com/andersfylling/disgord"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
)

func findChannel(channel string) bool {
	for _, c := range config.Channels {
		if c == channel {
			return true
		}
	}

	return false
}

func handleMsg(s disgord.Session, data *disgord.MessageCreate) {
	m := data.Message

	if findChannel(m.ChannelID.String()) {
		return
	}

	err := s.Channel(m.ChannelID).TriggerTypingIndicator()
	common.CheckErr(err, "trigger typing")

	if len(m.Attachments) > 0 {
		for _, a := range m.Attachments {
			uploadImg(s, m, a.URL)
		}
	} else {
		uploadImg(s, m, m.Content)
	}

	deleteMsg(s, m, 2*time.Second)
}
