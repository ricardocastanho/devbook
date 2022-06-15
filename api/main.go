package main

import (
	// "crypto/rand"
	// "encoding/base64"
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

// func init() {
// 	key := make([]byte, 64)

// 	_, err := rand.Read(key)

// 	if err != nil {
// 		log.Fatal("Error: ", err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)

// 	fmt.Println("Generated key: ", stringBase64)
// }

func main() {
	config.LoadEnv()

	r := router.BuildRouter()

	fmt.Printf("Listening on port :%s 🚀\n", config.APIPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.APIPort), r))
}
