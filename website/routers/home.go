package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
