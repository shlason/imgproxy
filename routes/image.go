package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shlason/imgproxy/controllers"
)

func RegisteImageRoutes(r *gin.RouterGroup) {
	r.GET("/image", controllers.GetImagesByQS)
}
