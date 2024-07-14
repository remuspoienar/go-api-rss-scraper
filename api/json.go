package api

import (
	"encoding/json"
	"net/http"
)

func sendJson(w http.ResponseWriter, code int, payload any) error {
	response, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return nil
}

func sendError(w http.ResponseWriter, code int, message string) error {
	return sendJson(w, code, map[string]string{"error": message})
}
