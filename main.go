package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/suryab-21/golang-template/app/router"
	"github.com/suryab-21/golang-template/app/service"
)

// @title Golang Template
// @version 1.0
// @description For test purpose
// @host localhost:8080
// @BasePath /
// @securitydefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Bearer token
func main() {
	service.InitDB()
	service.InitCache()

	fmt.Println("Server listening on port :" + os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(`:%s`, os.Getenv("PORT")), router.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
