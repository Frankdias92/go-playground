package apirest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func sendJSON(w http.ResponseWriter, resp Response, status int) {
	data, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("error to do marshal of json:", err)
		sendJSON(w, Response{Error: "something went wrong"}, http.StatusInternalServerError)
	}

	w.WriteHeader(status)
	if _, err := w.Write(data); err != nil {
		fmt.Println("error to send response", err)
	}
}

type User struct {
	Username string
	ID       int64 `json:"id,string"`
	Role     string
	Password string `json:"-"`
}

var db = map[int64]User{
	1: {
		ID:       1,
		Username: "admin",
		Password: "admin",
		Role:     "admin",
	},
}

func GetJson() {
	opts := &slog.HandlerOptions{
		AddSource:   false,
		Level:       nil,
		ReplaceAttr: nil,
	}
	l := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	slog.SetDefault(l)

	l = l.With(slog.Group("app_info", slog.String("version:", "1.0.0")))

	l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		"Info Message",
		slog.String("key", "value"),
		slog.Group("group message",
			slog.String("groupKey", "groupValue"),
			slog.String("groupKey2", "groupValue2"),
		),
		slog.Int("status", http.StatusOK),
	)

	r := chi.NewMux()

	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Use(jsonMiddleware)
		r.Get("/users/{id:[0-9]+}", handleGetUsers)
		r.Post("/users", handlePostUsers)
	})

	if err := http.ListenAndServe(":8085", r); err != nil {
		panic(err)
	}
}

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func handleGetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := chi.URLParam(r, "id")
	id, _ := strconv.ParseInt(idStr, 10, 64)

	user, ok := db[id]
	if !ok {
		sendJSON(w, Response{Error: "user not found"}, http.StatusNotFound)
		return
	}
	sendJSON(w, Response{Data: user}, http.StatusOK)
}

func handlePostUsers(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10000)
	data, err := io.ReadAll(r.Body)

	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			sendJSON(w, Response{Error: "body too large"}, http.StatusRequestEntityTooLarge)
			return
		}

		fmt.Println(err)
		return
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		sendJSON(w, Response{Error: "invalid body"}, http.StatusUnprocessableEntity)
		return
	}

	db[user.ID] = user

	w.WriteHeader(http.StatusCreated)

}
