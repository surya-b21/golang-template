package router

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/suryab-21/sigmatech-test/app/controllers/myclient"
)

func InitRoutes() http.Handler {
	router := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	router.HandleFunc("GET /client", myclient.GetMyClient)
	router.HandleFunc("POST /client", myclient.PostMyClient)
	router.HandleFunc("GET /client/{id}", myclient.GetByIdMyClient)
	router.HandleFunc("PUT /client/{id}", myclient.PutMyClient)
	router.HandleFunc("DELETE /client/{id}", myclient.DeleteMyClient)

	return cors.Handler(router)
}
