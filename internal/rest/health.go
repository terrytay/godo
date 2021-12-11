package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/terrytay/godo/internal/service"
)

type HealthService interface {
	Get(ctx context.Context) *service.HealthResponse
}

type HealthHandler struct {
	svc HealthService
}

func NewHealthHandler(svc HealthService) *HealthHandler {
	return &HealthHandler{
		svc: svc,
	}
}

func (h *HealthHandler) Register(r *mux.Router) {
	r.HandleFunc("/health", h.get).Methods(http.MethodGet)
}

func (h *HealthHandler) get(rw http.ResponseWriter, r *http.Request) {
	response := h.svc.Get(r.Context())
	rw.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(rw)
	e.Encode(response)
}
