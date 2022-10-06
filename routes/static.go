package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shlason/imgproxy/configs"
)

func RegisteStaticContentRoutes(r *gin.Engine) {
	r.Static("/static", fmt.Sprintf("%s/static/", configs.Server.FeWrokDir))
	r.StaticFile("/", fmt.Sprintf("%s/index.html", configs.Server.FeWrokDir))
}
