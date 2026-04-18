package views

import (
	"encoding/json"
	"net/http"
)

type HealthzResponse struct {
	Status string `json:"status" example:"ok"`
}

// GetHealth godoc
//
//	@Summary		Check API health
//	@Description	Returns the status of the API
//	@Tags			health healthz
//	@Produce		json
//	@Success		200	{object}	HealthzResponse
//	@Router			/healthz [get]
func GetHealthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(HealthzResponse{Status: "ok"})
}
