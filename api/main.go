package main

import (
    "fmt"
    "log"
    "net/http"

    "api/src/router"
)

func main() {
    fmt.Println("Listening on port :3000 🚀")

    r := router.BuildRoutes()

    log.Fatal(http.ListenAndServe(":3000", r))
}
