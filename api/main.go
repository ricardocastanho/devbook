package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	config.LoadEnv()

	r := router.BuildRouter()

	fmt.Printf("Listening on port :%s 🚀", config.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.APIPort), r))
}
