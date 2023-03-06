package handler

import (
	"fmt"
	"github.com/cucumberjaye/softline/internal/models"
	"github.com/go-chi/render"
	"github.com/go-playground/validator"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	var input models.RegisterUser

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		log.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(&input)
	if err != nil {
		log.Error().Err(err)
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	id, err := h.authService.CreateUser(input)
	if err != nil {
		log.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]int{
		"id": id,
	})
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input models.LoginUser

	err := render.DecodeJSON(r.Body, &input)
	if err != nil {
		log.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = validator.New().Struct(&input)
	if err != nil {
		log.Error().Err(err)
		http.Error(w, fmt.Sprintf("invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	token, err := h.authService.GenerateToken(input)
	if err != nil {
		log.Error().Err(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, map[string]string{
		"token": token,
	})
}
