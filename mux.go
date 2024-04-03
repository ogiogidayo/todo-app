package main

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/go-playground/validator/v10"
	"github.com/ogiogidayo/todo-app/clock"
	"github.com/ogiogidayo/todo-app/config"
	"github.com/ogiogidayo/todo-app/database"
	"github.com/ogiogidayo/todo-app/handler"
	"github.com/ogiogidayo/todo-app/services"
	"net/http"
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
	r := database.Repository{Clocker: clock.RealClocker{}}
	at := &handler.AddTask{
		Services:  &services.AddTask{DB: db, Repo: &r},
		Validator: v,
	}
	mux.Post("/tasks", at.ServeHTTP)
	lt := &handler.ListTask{
		Sevices: &services.ListTask{DB: db, Repo: &r},
	}
	mux.Get("/tasks", lt.ServeHTTP)
	return mux, cleanup, err
}
