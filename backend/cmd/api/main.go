package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"sezzle-calculator/backend/internal/calculator"
	"sezzle-calculator/backend/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	r := chi.NewRouter()

	// Middlewares estándar
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Configuración CORS permisiva para desarrollo local y frontend
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Inicialización de componentes
	calcService := calculator.NewService()
	calcHandler := handler.NewCalculatorHandler(calcService)

	// Rutas
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	r.Post("/api/v1/calculate", calcHandler.Calculate)

	fmt.Printf("Servidor corriendo en el puerto %s...\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
