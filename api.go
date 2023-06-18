package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

type apiServer struct {
	listenAddr string
}

func newAPIServer(listenAddr string) *apiServer {
	return &apiServer{
		listenAddr: listenAddr,
	}
}

func (s *apiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccout))

	log.Println("JSON API server running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *apiServer) handleAccout(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccout(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateAccout(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteAccout(w, r)
	}

	return fmt.Errorf("method not allowed %s", r.Method)
}

func (s *apiServer) handleGetAccout(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *apiServer) handleCreateAccout(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *apiServer) handleDeleteAccout(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *apiServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
