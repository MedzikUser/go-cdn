package imgur

import (
	"strings"

	"gitlab.com/gaming0skar123/go/cdn/config"
	"gitlab.com/gaming0skar123/go/modules/imgur"
)

func UploadFromURL(url string) (*imgur.ImageInfo, error) {
	i, _, err := clientImgur.UploadImageFromURL(url, config.Imgur_Album_ID)

	if err == nil {
		i.Data.Link = strings.Replace(i.Data.Link, "i.imgur.com", config.Proxy_Domain, -1)

		go insert(i.Data)
	}

	if i != nil {
		return i.Data, err
	}

	return nil, err
}

func Upload(b []byte) (*imgur.ImageInfo, error) {
	i, _, err := clientImgur.UploadImage(string(b[:]), "file", config.Imgur_Album_ID)

	if err == nil {
		i.Data.Link = strings.Replace(i.Data.Link, "i.imgur.com", config.Proxy_Domain, -1)

		go insert(i.Data)
	}

	if i != nil {
		return i.Data, err
	}

	return nil, err
}
