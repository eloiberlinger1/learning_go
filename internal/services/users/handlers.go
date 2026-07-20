package users

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct {
	store *Store
}

func (h *Handler) RegisterRoutes(router chi.Router) {

	router.Post("/register", h.HandleRegister)
	router.Post("/login", h.HandleLogin)

}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

// HashPassword est une fonction utilitaire locale pour hacher le mot de passe avant stockage.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (h *Handler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var payload LoginUserPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if payload.Email == "" || payload.Password == "" {
		http.Error(w, "credentials should not be empty", http.StatusBadRequest)
		return
	}

	user, err := h.store.GetUserByEmail(r.Context(), payload.Email)
	if err != nil {
		http.Error(w, "Email ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password))
	if err != nil {
		http.Error(w, "Email ou mot de passe incorrect", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)

}

func (h *Handler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload RegisterUserPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "JSON invalide", http.StatusBadRequest)
		return
	}

	if payload.Email == "" || payload.Password == "" {
		http.Error(w, "L'email et le mot de passe sont obligatoires", http.StatusBadRequest)
		return
	}

	_, err := h.store.GetUserByEmail(r.Context(), payload.Email)
	if err == nil {
		http.Error(w, "Un utilisateur avec cet email existe déjà", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(payload.Password)
	if err != nil {
		http.Error(w, "Erreur interne de hachage", http.StatusInternalServerError)
		return
	}

	newUser := &User{
		Email:        payload.Email,
		PasswordHash: hashedPassword,
		FirstName:    payload.FirstName,
		LastName:     payload.LastName,
	}

	if err := h.store.CreateUser(r.Context(), newUser); err != nil {
		http.Error(w, "Erreur lors de la création de l'utilisateur en BDD", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}
