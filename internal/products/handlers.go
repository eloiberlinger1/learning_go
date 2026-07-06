package products

import (
	"ecom-local/internal/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	products, err := h.service.ListProducts(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, products)
}

func (h *handler) ListProductById(w http.ResponseWriter, r *http.Request) {
	productID := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(productID, 10, 64)
	if err != nil {
		http.Error(w, "ID de produit invalide", http.StatusBadRequest)
		return
	}

	product, err := h.service.ListProductById(r.Context(), id)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, product)
}
