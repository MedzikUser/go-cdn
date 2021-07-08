package imgur

import (
	"net/http"

	"gitlab.com/gaming0skar123/go/cdn/config"
	"gitlab.com/gaming0skar123/go/modules/imgur"
)

var clientImgur = new(imgur.Client)

func init() {
	clientImgur.HTTPClient = new(http.Client)
	clientImgur.ImgurClientID = config.Imgur_Client_ID
}
