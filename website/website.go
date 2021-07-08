package website

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/gaming0skar123/go/cdn/website/routers"
)

func Server() {
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.MaxMultipartMemory = 1

	r.GET("/", routers.Home)
	r.POST("/upload", routers.Upload)

	err := r.Run(":8080")
	if err != nil {
		panic("Failed Starting GIN server: " + err.Error())
	}
}
