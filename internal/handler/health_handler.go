package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/joserafaelSH/go-api/internal/telemetry"
	"go.opentelemetry.io/otel/codes"
)

func HealthHandler(ctx context.Context, otel *telemetry.Jaeger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		oplog := httplog.LogEntry(r.Context())
		_, span := otel.Start(ctx, "healthHandler")
		defer span.End()
		var result struct {
			Status string `json:"status"`
		}
		result.Status = "OK"
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(result); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			oplog.Error().Msg("failed to encode response")
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
			return
		}

	}
}
