package middleware

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Contetnt-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
