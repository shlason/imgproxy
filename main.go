package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/shlason/imgproxy/configs"
	"github.com/shlason/imgproxy/docs"
	"github.com/shlason/imgproxy/routes"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

// @title           Image-Proxy Example API
// @version         1.0
// @description     This is a sample server celler server.

// @contact.name   sidesideeffect.io
// @contact.url    https://github.com/shlason/imgproxy
// @contact.email  nocvi111@gmail.com

// @license.name  MIT
// @license.url   https://github.com/shlason/imgproxy/blob/main/LICENSE

// @host      imgproxy.sidesideeffect.io
// @BasePath  /api
func main() {
	var g errgroup.Group

	localPtr := flag.Bool("local", false, "Running in local?")
	flag.Parse()

	r := gin.Default()
	r.Use(cors.Default())
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"https://*.imgproxy.sidesideeffect.io"},
	// 	AllowMethods: []string{"GET", "OPTIONS"},
	// 	AllowHeaders: []string{"Origin"},
	// }))

	apiRoute := r.Group("/api")

	routes.RegisteStaticContentRoutes(r)
	routes.RegisteImageRoutes(apiRoute)

	docs.SwaggerInfo.Schemes = []string{"https"}

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if *localPtr {
		r.Run()
		return
	}

	g.Go(func() error {
		return http.ListenAndServe(":http", http.RedirectHandler(fmt.Sprintf("https://%s", configs.Server.Host), http.StatusSeeOther))
	})
	g.Go(func() error {
		return http.Serve(autocert.NewListener(configs.Server.Host), r)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
