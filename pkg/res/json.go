package res

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, body interface{}) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(body)
}