package res

import (
	"encoding/json"
	"net/http"
)

type Map map[string]interface{}

type Error struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func RenderJSON(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&body)
}

func ErrorJSON(w http.ResponseWriter, status int, message interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	errorResponse := Error{
		Code:    status,
		Message: message,
	}

	json.NewEncoder(w).Encode(&errorResponse)
}
