package imgur

import (
	"gitlab.com/gaming0skar123/go/cdn/database"
	"gitlab.com/gaming0skar123/go/modules/imgur"
)

func insert(i *imgur.ImageInfo) {
	database.InsertImgur(&database.Imgur{
		ID:      i.IDExt,
		DelHash: i.Deletehash,
	})
}
