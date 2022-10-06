package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shlason/imgproxy/configs"
)

func RegisteStaticContentRoutes(r *gin.Engine) {
	r.StaticFile("/", fmt.Sprintf("%s/index.html", configs.Server.FeWrokDir))
}
