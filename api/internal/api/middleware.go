package api

import (
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/render"
)

func APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := os.Getenv("API_KEY")
		if apiKey == "" {
			next.ServeHTTP(w, r)
			return
		}

		requestKey := r.Header.Get("X-API-Key")
		if requestKey == "" {
			render.JSON(w, r, map[string]string{"error": "API key is required"})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !strings.EqualFold(requestKey, apiKey) {
			render.JSON(w, r, map[string]string{"error": "Invalid API key"})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
