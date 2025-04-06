package middleware

import (
	"log/slog"
	"net/http"
)

// simpleHeaderMiddleware adds a static header to the response.
// It takes the next handler in the chain and returns a handler.
func SimpleHeaderMiddleware(slogger *slog.Logger, next http.Handler) http.Handler {
	// http.HandlerFunc is an adapter, allowing ordinary functions to be used as handlers.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// --- Logic before the next handler ---
		// Set a header before passing the request down.
		// Note: If the next handler also sets this header, the last one wins
		// unless you use .Add() instead of .Set() for multi-value headers.
		w.Header().Set("X-Simple-Middleware", "Applied")

		// --- Call the next handler ---
		// This executes the handler that was wrapped by this middleware.
		next.ServeHTTP(w, r)

		// --- Logic after the next handler (optional) ---
		// We could do something here after the main handler has finished,
		// but for this example, we don't need to.
	})
}
