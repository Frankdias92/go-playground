package myfirtsapigo

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		begin := time.Now()
		next.ServeHTTP(w, r)
		fmt.Println(r.URL.String(), r.Method, time.Since(begin))
	})
}

func Server() {
	mux := http.NewServeMux()

	mux.HandleFunc(
		"/api/user/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("id")
			fmt.Println(id)
			fmt.Fprintf(w, "hello, world\n")
		},
	)

	srv := &http.Server{
		Addr:                         ":8085",
		Handler:                      Log(mux),
		DisableGeneralOptionsHandler: false,
		ReadTimeout:                  10 * time.Second,
		ReadHeaderTimeout:            10 * time.Second,
		WriteTimeout:                 10 * time.Second,
		IdleTimeout:                  1 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
	srv.ListenAndServe()
}
