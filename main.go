package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/suryab-21/sigmatech-test/app/router"
	"github.com/suryab-21/sigmatech-test/app/service"
)

func main() {
	service.InitDB()
	service.InitCache()

	fmt.Println("Server listening on port :" + os.Getenv("PORT"))
	if err := http.ListenAndServe(fmt.Sprintf(`:%s`, os.Getenv("PORT")), router.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
