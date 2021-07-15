package discord

import (
	"time"

	"github.com/andersfylling/disgord"
	"gitlab.com/gaming0skar123/go/cdn/common"
)

func deleteMsg(s disgord.Session, m *disgord.Message, t time.Duration) {
	go func() {
		time.Sleep(t)
		err := s.Channel(m.ChannelID).Message(m.ID).Delete()
		common.CheckErr(err, "delete msg")
	}()
}
