package middleware

import (
	"encoding/json"
	"net/http"
)

type JSONHandlerFuncWithWriter func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

func WriteJSON(h JSONHandlerFuncWithWriter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, status, err := h(w, r)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)

		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}

		json.NewEncoder(w).Encode(data)
	}
}
