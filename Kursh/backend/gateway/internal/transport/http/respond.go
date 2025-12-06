package http

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, code int, msg string) {
	writeJSON(w, code, map[string]string{"error": msg})
}

func extractID(path, prefix string) (string, bool) {
	if len(path) <= len(prefix) {
		return "", false
	}
	id := path[len(prefix):]
	if id == "" {
		return "", false
	}
	return id, true
}
