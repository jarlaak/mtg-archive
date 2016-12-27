package server

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, output interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}
