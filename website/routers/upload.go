package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/cdn/common"
	"gitlab.com/gaming0skar123/go/cdn/discord"
	"gitlab.com/gaming0skar123/go/cdn/imgur"
)

func Upload(c *gin.Context) {
	conLen := c.Request.ContentLength

	if conLen > 10485956 /* 10 MB */ {
		c.String(http.StatusRequestEntityTooLarge, "Request Too Large")
		return
	}

	if conLen <= 0 {
		c.String(http.StatusBadRequest, "Request Length is Unknown or Empty")
		return
	}

	file, err := c.FormFile("file")
	if common.CheckErr(err, "form file") {
		c.String(http.StatusBadRequest, "Error Form File")
		return
	}

	f, err := file.Open()
	if common.CheckErr(err, "open file") {
		c.String(http.StatusBadRequest, "Error Opening File")
		return
	}
	defer f.Close()

	size := file.Size

	b := make([]byte, size)

	n, err := f.Read(b)
	if err != nil || int64(n) != size {
		c.String(http.StatusBadRequest, "Error Reading File")
		return
	}

	i, err := imgur.Upload(b)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	go discord.API(i.Link)

	c.String(http.StatusOK, "Uploaded!")
}
