package handler

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog"
	"github.com/joserafaelSH/go-api/config/logger"
	"github.com/joserafaelSH/go-api/internal/telemetry"
)

func InitHttpHandler() {
	ctx := context.Background()

	otel, err := telemetry.NewJaeger(ctx, "go-api")
	if err != nil {
		logger.Log.Writter.Panic().Msg(err.Error())
	}
	defer otel.Shutdown(ctx)

	router := chi.NewRouter()
	router.Use(httplog.RequestLogger(logger.Log.Writter))

	router.Get("/api/v1/health", HealthHandler(ctx, otel))

	http.Handle("/", router)

	srv := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      http.DefaultServeMux,
	}
	err = srv.ListenAndServe()
	if err != nil {
		logger.Log.Writter.Panic().Msg(err.Error())
	}
}
