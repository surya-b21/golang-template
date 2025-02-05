package router

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/suryab-21/golang-template/app/controllers/auth"
	"github.com/suryab-21/golang-template/app/controllers/user"
	"github.com/suryab-21/golang-template/app/middleware"
)

func InitRoutes() http.Handler {
	router := http.NewServeMux()

	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	authorizedMiddleware := middleware.MiddlewareStack(
		middleware.ClaimToken,
	)

	router.Handle("/api/", http.StripPrefix("/api", authorizedMiddleware(authorizedRoute())))
	router.HandleFunc("POST /api/sign-up", auth.SignUp)
	router.HandleFunc("POST /api/sign-in", auth.SignIn)

	return cors.Handler(router)
}

func authorizedRoute() *http.ServeMux {
	authorizedRoute := http.NewServeMux()

	authorizedRoute.HandleFunc("GET /users/me", user.GetUserInfo)

	return authorizedRoute
}
