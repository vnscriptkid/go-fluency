package main

import (
	"bytes"
	"net/http"

	"github.com/google/uuid"
)

type Store interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

type MemStore struct {
	store map[string]string
}

func (s *MemStore) Get(key string) (string, bool) {
	val, ok := s.store[key]
	return val, ok
}

func (s *MemStore) Set(key, value string) {
	s.store[key] = value
}

type IdempotentHandler struct {
	store Store
	next  http.Handler
}

type responseWriter struct {
	http.ResponseWriter
	body   *bytes.Buffer
	status int
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *responseWriter) Write(b []byte) (int, error) {
	if w.body == nil {
		w.body = &bytes.Buffer{}
	}
	return w.body.Write(b)
}

func (h *IdempotentHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("X-Idempotency-Key")
	if id == "" {
		id = uuid.New().String()
	}

	if _, ok := h.store.Get(id); ok {
		w.WriteHeader(400)
		w.Write([]byte("Repeated action"))
		return
	}

	w = &responseWriter{ResponseWriter: w}
	h.next.ServeHTTP(w, r)

	h.store.Set(id, w.(*responseWriter).body.String())
}

func main() {
	store := &MemStore{store: map[string]string{}}

	http.Handle("/", &IdempotentHandler{
		store: store,
		next: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Got through"))
		}),
	})

	http.ListenAndServe(":8080", nil)
}
