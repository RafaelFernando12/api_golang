package controller

import (
	"net/http"

	"github.com/RafaelFernando12/api-golang/pkg/log"
	"github.com/gin-gonic/gin"
)

type handler struct {
	log log.MultiLogger
}

func NewHandler(
	log log.MultiLogger,
	origin string) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	return router.Handler()
}
