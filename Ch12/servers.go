package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

func slowServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		ctx := r.Context()
		select {
		case <-ctx.Done():
			fmt.Println("server shut down")
			return
		case <-time.After(6 * time.Second):
			w.Write([]byte("Slow response"))
		}
	}))
	return s
}

func fastServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter,
		r *http.Request) {
		if r.URL.Query().Get("error") == "true" {
			w.Write([]byte("error"))
			return
		}
		w.Write([]byte("ok"))
	}))
	return s
}
