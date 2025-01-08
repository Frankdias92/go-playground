package apiproject

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type getShortenenURLResponse struct {
	FullURL string `json:"url"`
}

func handleGetShortenendURL(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		fullURL, ok := db[code]
		if !ok {
			sendJson(w, apiRespose{Error: "code not found"}, http.StatusNotFound)
			return
		}

		sendJson(w, apiRespose{Data: getShortenenURLResponse{FullURL: fullURL}}, http.StatusOK)
	}
}
