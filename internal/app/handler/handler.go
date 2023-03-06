package handler

import (
	"github.com/cucumberjaye/softline/internal/models"
	"github.com/go-chi/chi/v5"
)

type AuthService interface {
	CreateUser(user models.RegisterUser) (int, error)
	GenerateToken(user models.LoginUser) (string, error)
}

type Handler struct {
	authService AuthService
}

func New(authService AuthService) *Handler {
	return &Handler{
		authService: authService,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	return r
}
