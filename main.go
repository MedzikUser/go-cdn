package main

import (
	"github.com/jpillora/opts"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
	_ "gitlab.com/gaming0skar123/go/cdn/database"
	"gitlab.com/gaming0skar123/go/cdn/discord"
	"gitlab.com/gaming0skar123/go/cdn/update"
	"gitlab.com/gaming0skar123/go/cdn/website"
)

type CMDOptions struct {
	Update bool `opts:"help=automatic updates"`
}

var log = common.Log

func main() {
	cmd := CMDOptions{
		Update: true,
	}

	opts.Parse(&cmd)

	log.Info("You're using verion: ", config.Version)

	if cmd.Update {
		go update.Updater()
	} else {
		log.Warn("Auto Update -> Disabled")
	}

	go website.Server()
	discord.Bot()
}
