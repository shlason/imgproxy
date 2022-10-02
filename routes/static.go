package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisteStaticContentRoutes(r *gin.Engine) {
	r.StaticFile("/", "/react_app/actions-runner/image_proxy/Image-Proxy/Image-Proxy/build/index.html")
}
