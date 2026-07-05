package products

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts() {
	// call service
	// return json to using api.go functions in order to genreate http repsonse
}
