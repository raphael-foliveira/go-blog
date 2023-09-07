package res

import "net/http"

func SendStatus(w http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	return nil
}
