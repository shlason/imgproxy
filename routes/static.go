package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shlason/imgproxy/configs"
)

func RegisteStaticContentRoutes(r *gin.Engine) {
	r.StaticFile("/", fmt.Sprintf("%s%s", configs.Server.FeWrokDir, "index.html"))
}
