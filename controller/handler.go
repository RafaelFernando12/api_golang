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

	// handler := &handler{
	// 	log: log,
	// }

	router.RemoveExtraSlash = true
	router.RedirectTrailingSlash = true

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	return router.Handler()
}
