package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) logMiddleware(c *gin.Context) {
	started := time.Now()
	defer func() {
		elapsedTime := time.Since(started)
		h.log.Info().Printf("%s %s %s %d", c.Request.Method, c.Request.URL, c.RemoteIP(), elapsedTime)
	}()

	c.Next()
}
