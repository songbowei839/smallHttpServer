package zapis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"pkg/config"
)

func Bind(r *gin.Engine) {
	r.POST("modez", postMode)
}

func postMode(c *gin.Context) {
	debug := config.SwitchMode()
	if debug {
		c.String(http.StatusOK, "debug\n")
	} else {
		c.String(http.StatusOK, "release\n")
	}
}
