package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shlason/imgproxy/configs"
	"github.com/shlason/imgproxy/routes"
)

func main() {
	// var g errgroup.Group

	r := gin.Default()

	apiRoute := r.Group("/api")

	routes.RegisteStaticContentRoutes(r)
	routes.RegisteImageRoutes(apiRoute)

	err := http.ListenAndServeTLS(fmt.Sprintf(":%s", configs.Server.Port), "../../../../cert/server.pem", "../../../../cert/server.key", r)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	// g.Go(func() error {
	// 	return http.ListenAndServe(":http", http.RedirectHandler(fmt.Sprintf("https://%s", configs.Server.Host), http.StatusSeeOther))
	// })
	// g.Go(func() error {
	// 	return http.Serve(autocert.NewListener(configs.Server.Host), r)
	// })

	// if err := g.Wait(); err != nil {
	// 	log.Fatal(err)
	// }
}
