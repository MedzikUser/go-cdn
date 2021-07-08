package update

import (
	"fmt"
	"net/http"
	"os"

	"github.com/inconshreveable/go-update"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/config"
)

func Update() {
	fmt.Println("Updating...")

	url := config.Latest_Binary

	resp, err := http.Get(url)
	if common.CheckErr(err, "downloading latest binary") {
		return
	}

	defer resp.Body.Close()

	err = update.Apply(resp.Body, update.Options{})
	if common.CheckErr(err, "self-update binary") {
		return
	}

	fmt.Println("Updated!")

	defer os.Exit(0)
}
