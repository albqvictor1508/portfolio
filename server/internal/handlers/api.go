package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"	
	"github.com/portfolio-api/server/internal/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use()
}
