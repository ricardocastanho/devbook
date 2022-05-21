package presenters

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		log.Fatal(err)
	}
}
