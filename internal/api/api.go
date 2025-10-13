package api

import (
	"github.com/Luiz-Hen-Reis/go-book-library/internal/services"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router *chi.Mux
	AuthorService services.AuthorService
}