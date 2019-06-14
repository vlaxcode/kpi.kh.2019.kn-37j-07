package api

import (
	"github.com/jaxmef/booklover-backend/workers/api/handler"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jaxmef/booklover-backend/config"
	"github.com/lancer-kit/armory/api"
	"github.com/lancer-kit/armory/api/render"
	"github.com/lancer-kit/armory/log"
	"github.com/sirupsen/logrus"
)

func Server() *api.Server {
	return &api.Server{
		Name:      "api-server",
		GetRouter: GetRouter,
		GetConfig: func() api.Config {
			return config.Config().Api
		},
	}
}

func GetRouter(logger *logrus.Entry, config api.Config) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(log.NewRequestLogger(logger.Logger))

	if config.EnableCORS {
		corsHandler := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "jwt", "X-UID"},
			ExposedHeaders:   []string{"Link", "Content-Length"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		})
		r.Use(corsHandler.Handler)
	}

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	if config.ApiRequestTimeout > 0 {
		t := time.Duration(config.ApiRequestTimeout)
		r.Use(middleware.Timeout(t * time.Second))
	}

	r.Route("/v1", func(r chi.Router) {
		r.Post("/register", handler.RegisterAcc)
		r.Get("/last-registered", handler.GetLastAccount)
	})

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		render.ResultNotFound.Render(w)
	})

	return r
}
