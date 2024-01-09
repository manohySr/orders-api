package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/manohySr/orders-api/handler"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrderRoutes)

	return router
}

func loadOrderRoutes(router chi.Router) {
	orderRoute := &handler.Order{}

	router.Get("/", orderRoute.List)
	router.Post("/", orderRoute.Create)
	router.Get("/{id}", orderRoute.GetByID)
	router.Patch("/{id}", orderRoute.UpdateByID)
	router.Delete("/{id}", orderRoute.DeleteByID)
}
