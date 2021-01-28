package router

import (
	"github.com/go-chi/chi"
	"github.com/pished/river-styx/api"
	"github.com/pished/river-styx/requestlog"
	"github.com/pished/river-styx/router/middleware"
)

func New(a *api.Api) *chi.Mux {
	l := a.Logger()

	r := chi.NewRouter()
	r.Method("GET", "/", requestlog.NewHandler(a.HandleIndex, l))

	r.Route("/v1/", func(r chi.Router) {
		r.Use(middleware.ContentTypeJson)

		// Routes for water
		r.Method("GET", "/water", requestlog.NewHandler(a.HandleListWater, l))
		r.Method("POST", "/water", requestlog.NewHandler(a.HandleAddWater, l))
		r.Method("GET", "/water/{id}", requestlog.NewHandler(a.HandleGetWater, l))
		r.Method("PUT", "/water/{id}", requestlog.NewHandler(a.HandleUpdateWater, l))
		r.Method("DELETE", "/water/{id}", requestlog.NewHandler(a.HandleDeleteWater, l))
	})

	return r
}
