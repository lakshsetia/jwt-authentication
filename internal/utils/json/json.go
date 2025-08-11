package json

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, value any, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(value); err != nil {
		http.Error(w, "failed to encode response json: " + err.Error(), http.StatusInternalServerError)
	}
}

func ReadJSON(r *http.Request, value any) error {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(value); err != nil {
		return err
	}
	return nil
}