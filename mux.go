package main

import (
	"context"
	"github.com/ogiogidayo/todo-app/auth"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/ogiogidayo/todo-app/clock"
	"github.com/ogiogidayo/todo-app/config"
	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/handler"
	"github.com/ogiogidayo/todo-app/services"
)

func NewMux(ctx context.Context, cfg *config.Config) (http.Handler, func(), error) {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	db, cleanup, err := database.New(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	clocker := clock.RealClocker{}
	r := database.Repository{Clocker: clocker}
	rcli, err := database.NewKVS(ctx, cfg)
	if err != nil {
		return nil, cleanup, err
	}
	jwter, err := auth.NewJWTer(rcli, clocker)
	if err != nil {
		return nil, cleanup, err
	}
	ru := &handler.RegisterUser{
		Service:   &services.RegisterUser{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/register", ru.ServeHTTP)
	l := &handler.Login{
		Service: &services.Login{
			DB:             db,
			Repo:           &r,
			TokenGenerator: jwter,
		},
		Validatpr: v,
	}
	mux.Post("/login", l.ServeHTTP)
	at := &handler.AddTask{
		Services:  &services.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	lt := &handler.ListTask{
		Sevices: &services.ListTask{DB: db, Repo: &r},
	}
	mux.Route("/tasks", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter))
		r.Post("/", at.ServeHTTP)
		r.Get("/", lt.ServeHTTP)
	})
	mux.Route("/admin", func(r chi.Router) {
		r.Use(handler.AuthMiddleware(jwter), handler.AdminMiddleware)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			_, _ = w.Write([]byte(`{"message": "admin only"}`))
		})
	})
	return mux, cleanup, err
}
