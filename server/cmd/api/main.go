package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/portfolio-api/server/internal/handlers"
)

func main() {
	log.SetReporterCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá, Mundo!"))
	})

	log.Println("HTTP Server running on 8080")
	err := http.ListenAndServe(":8080", r);

	err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
