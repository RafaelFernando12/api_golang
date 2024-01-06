package controller

import (
	"net/http"

	"github.com/RafaelFernando12/api-golang/domain"
	"github.com/RafaelFernando12/api-golang/pkg/log"
	"github.com/gin-gonic/gin"
)

type handler struct {
	userService domain.UserService
	log         log.MultiLogger
}

func NewHandler(
	userService domain.UserService,
	log log.MultiLogger,
	origin string) http.Handler {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	handler := &handler{
		userService: userService,
		log:         log,
	}

	router.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	v1 := router.Group("/v1", handler.logMiddleware)
	user := v1.Group("/user")
	user.POST("", handler.createUser)
	return router.Handler()
}
