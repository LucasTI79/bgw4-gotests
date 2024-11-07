package utils

import (
	"context"
	"database/sql"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/lucasti79/bgw4-put-patch-delete/internal/server"
)

func CreateServerTest(t *testing.T, db *sql.DB) *chi.Mux {
	t.Helper()

	return server.CreateServer(db)
}

func WithUrlParam(t *testing.T, r *http.Request, key, value string) *http.Request {
	t.Helper()
	chiCtx := chi.NewRouteContext()
	chiCtx.URLParams.Add(key, value)
	req := r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, chiCtx))

	return req
}

// rctx := chi.NewRouteContext()
// rctx.URLParams.Add("key", "value")

// r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
