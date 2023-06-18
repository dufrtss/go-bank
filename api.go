package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	return json.NewEncoder(w).Encode(v)
}

type APIFunc func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string
}

func makeHTTPHandleFunc(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func newAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccout))

	log.Println("JSON API server running on port", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccout(w http.ResponseWriter, r *http.Request) error {
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

func (s *APIServer) handleGetAccout(w http.ResponseWriter, r *http.Request) error {
	account := NewAccount("Eduardo", "Freitas")

	return WriteJSON(w, http.StatusOK, account)
}

func (s *APIServer) handleCreateAccout(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccout(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
