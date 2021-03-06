package update

import (
	"time"

	"gitlab.com/gaming0skar123/go/cdn/config"
)

func CheckUpdates() {
	// Check on start
	Update()

	ticker := time.NewTicker(config.Latest_Version_Check)

	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			Update()

		case <-quit:
			ticker.Stop()

			return
		}
	}
}
