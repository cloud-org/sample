package main

import (
	"log"
	"net/http"
	"time"
)

type MiddleWare struct {
}

func (m MiddleWare) LoggingHandler(next http.Handler) http.Handler {
	var (
		t1 time.Time
		t2 time.Time
		fn func(w http.ResponseWriter, r *http.Request)
	)
	fn = func(w http.ResponseWriter, r *http.Request) {
		t1 = time.Now()
		next.ServeHTTP(w, r)
		t2 = time.Now()
		log.Println(r.Method, r.URL.String(), t2.Sub(t1))
	}

	return http.HandlerFunc(fn)
}

func (m MiddleWare) RecoverHandler(next http.Handler) http.Handler {
	var (
		fn  func(w http.ResponseWriter, r *http.Request)
		err interface{}
	)
	fn = func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err = recover(); err != nil {
				log.Println("recover from panic", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)

}
