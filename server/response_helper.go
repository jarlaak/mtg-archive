package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendError(w http.ResponseWriter, errorCode int, message string) {
	w.WriteHeader(errorCode)
	fmt.Fprintf(w, message)
}

func SendJSON(w http.ResponseWriter, output interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
