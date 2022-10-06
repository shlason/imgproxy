package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shlason/imgproxy/configs"
	"github.com/shlason/imgproxy/routes"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

func main() {
	var g errgroup.Group

	localPtr := flag.Bool("local", false, "Running in local?")
	flag.Parse()

	r := gin.Default()

	apiRoute := r.Group("/api")

	routes.RegisteStaticContentRoutes(r)
	routes.RegisteImageRoutes(apiRoute)

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
