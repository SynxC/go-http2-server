package main

import (
	"go-http2-server/config"
	_ "go-http2-server/docs"
	"go-http2-server/handlers"
	"go-http2-server/store"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	petStore := store.ConstructPetStore()

	http.HandleFunc(config.ApiURL, handlers.PostPets(petStore))
	http.HandleFunc(config.ApiURL+"/{id}", handlers.GetOrDeletePets(petStore))

	http.Handle("/swagger/*", httpSwagger.WrapHandler)

	http2Server := &http2.Server{}
	httpServer := http.Server{
		Addr: ":8080",
		Handler: h2c.NewHandler(http.DefaultServeMux, http2Server),
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}