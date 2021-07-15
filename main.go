package main

import (
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
	"gitlab.com/gaming0skar123/go/cdn/database"
	"gitlab.com/gaming0skar123/go/cdn/discord"
	"gitlab.com/gaming0skar123/go/cdn/website"
)

var log = common.Log

func main() {
	log.Info("You're using verion: ", config.Version)

	database.Connect()

	go website.Server()
	discord.Bot()
}
