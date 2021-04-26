package handlers

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	ReturnJson(w, http.StatusOK, "Healthy!", nil)
}
