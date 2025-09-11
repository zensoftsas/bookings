package main

import (
	"bookings/pkg/config"
	"bookings/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routers(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(NoSurf)
	r.Use(SessionLoad)

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)
	r.Get("/generals-quarters", handlers.Repo.Generals)
	r.Get("/majors-suite", handlers.Repo.Majors)

	r.Get("/search-availability", handlers.Repo.Availability)
	r.Post("/search-availability", handlers.Repo.PostAvailability)
	r.Get("/search-availability-json", handlers.Repo.AvailabilityJSON)
	
	r.Get("/contact", handlers.Repo.Contact)
	
	r.Get("/make-reservation", handlers.Repo.Reservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}
