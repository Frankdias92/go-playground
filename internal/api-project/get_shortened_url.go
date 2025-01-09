package apiproject

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/frankdias92/go-playground/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
)

type getShortenenURLResponse struct {
	FullURL string `json:"url"`
}

func handleGetShortenendURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fullURL, err := store.GetFullURL(r.Context(), code)
		if err != nil {
			if errors.Is(err, redis.Nil) {
				sendJson(w, apiRespose{Error: "code not found"}, http.StatusNotFound)
				return
			}
			slog.Error("failed to get code", "error", err)
			sendJson(w, apiRespose{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		sendJson(w, apiRespose{Data: getShortenenURLResponse{FullURL: fullURL}}, http.StatusOK)
	}
}
