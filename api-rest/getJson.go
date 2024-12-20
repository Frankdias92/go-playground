package apirest

import (
	"encoding/json"
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
	if ok {
		data, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		_, _ = w.Write(data)
	}
}

func handlePostUsers(w http.ResponseWriter, r *http.Request) {

}
