package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisteStaticContentRoutes(r *gin.Engine) {
	r.StaticFile("/", "./testing.html")
}
