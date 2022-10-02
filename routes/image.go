package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisteImageRoutes(r *gin.RouterGroup) {
	r.GET("/image", func(c *gin.Context) {
		c.Data(http.StatusOK, "application/json; charset=utf-8", []byte("Hi image"))
	})
}
