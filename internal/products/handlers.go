package products

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// call service
	// return json to using api.go functions in order to genreate http repsonse

	products := []any{"Hello", "Woorld", 5}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
