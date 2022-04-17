package res

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(w http.ResponseWriter, status int, body interface{}) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&body)
}
