package api

import "net/http"

func (c *Config) HealthCheck(w http.ResponseWriter, r *http.Request) {
	sendJson(w, 200, map[string]string{"status": "ok"})
}

func (c *Config) Err(w http.ResponseWriter, r *http.Request) {
	sendError(w, 500, "Internal Server Error")
}
