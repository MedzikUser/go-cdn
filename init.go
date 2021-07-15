package main

import (
	"github.com/jpillora/opts"
	"gitlab.com/gaming0skar123/go/cdn/update"
)

type CMDOptions struct {
	Update bool `opts:"help=automatic updates"`
}

func init() {
	cmd := CMDOptions{
		Update: true,
	}

	opts.Parse(&cmd)

	if cmd.Update {
		go update.CheckUpdates()
	} else {
		log.Warn("Auto Update -> Disabled")
	}
}
