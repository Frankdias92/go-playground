package apiprojectmovies

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/frankdias92/go-playground/omdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string
	Data  interface{}
}

func NewHandler() http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Get("/", handleSearchMovie())

	return r
}

func handleSearchMovie() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.URL.Query().Get("apikey")
		search := r.URL.Query().Get("s")
		res, err := omdb.Search(apiKey, search)

		if search == "" {
			sendJson(w, Response{Error: "search parameter 's' is required"}, http.StatusBadRequest)
			return
		}

		if apiKey == "" {
			sendJson(w, Response{Error: "apikey parameter is required"}, http.StatusUnauthorized)
			return
		}

		if err != nil {
			sendJson(w, Response{Error: "something wrong with omdb"}, http.StatusBadGateway)
			return
		}

		sendJson(w, Response{Data: res}, http.StatusOK)
	}
}

func sendJson(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to marshal json data", "error", err)
		sendJson(
			w,
			Response{Error: "something went wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		slog.Error("failed to write response to client", "error", err)
		return
	}
}
