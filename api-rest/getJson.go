package apirest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

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
		// http.Error(w, "user not found", http.StatusNotFound)
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write(([]byte(`{"error":"user not found"}`)))
		return
	}
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(data)
}

func handlePostUsers(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 10000)
	data, err := io.ReadAll(r.Body)

	if err != nil {
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			http.Error(w, "body too large", http.StatusRequestEntityTooLarge)
			return
		}

		fmt.Println(err)
		http.Error(w, "something went wrong", http.StatusInternalServerError)
		return
	}

	var user User
	if err := json.Unmarshal(data, &user); err != nil {
		http.Error(w, "invalid body", http.StatusUnprocessableEntity)
		return
	}

	db[user.ID] = user

	w.WriteHeader(http.StatusCreated)

}
