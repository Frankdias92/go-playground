package apiproject

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func sendJson(w http.ResponseWriter, resp Respose, status int) {
	w.Header().Set("Content-Type", "application/json")

	data, err := json.Marshal(resp)
	if err != nil {
		slog.Error("failed to mashal json data", "error", err)
		sendJson(
			w,
			Respose{Error: "something went wrong"},
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

func NewHandler(db map[string]string) http.Handler {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Post("/api/shorten", handlePost(db))
	r.Get("/{code}", handleGet(db))

	return r
}

type PostBody struct {
	URL string `json:"url"`
}

type Respose struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func genCode() string {
	const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const n = 8
	byts := make([]byte, n)

	// Inicializa a seed para gerar números aleatórios diferentes em cada execução
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		byts[i] = characters[rand.Intn(len(characters))]
	}

	return string(byts)
}

func handlePost(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body PostBody
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			sendJson(w, Respose{Error: "invalid body"}, http.StatusUnprocessableEntity)
			return
		}

		if _, err := url.Parse(body.URL); err != nil {
			sendJson(w, Respose{Error: "invalid url passed"}, http.StatusBadRequest)
		}

		code := genCode()
		db[code] = body.URL
		sendJson(w, Respose{Data: code}, http.StatusCreated)
	}
}

func handleGet(db map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")

		data, ok := db[code]

		if !ok {
			http.Error(w, "url not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, data, http.StatusPermanentRedirect)

	}
}
