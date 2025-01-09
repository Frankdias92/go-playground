package apiproject

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/frankdias92/go-playground/internal/store"
)

type shortenURLRequest struct {
	URL string `json:"url"`
}

type shortenURLResponse struct {
	Code string `json:"code"`
}

func handleShortenURL(store store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body shortenURLRequest

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJson(w, apiRespose{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJson(w, apiRespose{Error: "invalid url passed"}, http.StatusBadRequest)
			return
		}

		code, err := store.SaveShortendURL(r.Context(), body.URL)
		if err != nil {
			slog.Error("failed to create code", "error", err)
			sendJson(w, apiRespose{Error: "something went wrong"}, http.StatusInternalServerError)
			return
		}

		sendJson(w, apiRespose{Data: shortenURLResponse{Code: code}}, http.StatusCreated)
	}
}
