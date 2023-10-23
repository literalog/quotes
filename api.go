package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenAddr string
	store      Storage
}

func NewApiServer(listenAddr string, store Storage) *ApiServer {
	return &ApiServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/quotes", makeHttpHandler(s.handleGetQuotes)).Methods(http.MethodGet)

	http.ListenAndServe(s.listenAddr, router)
	log.Println("api server listening on", s.listenAddr)
}

func (s *ApiServer) handleQuotes(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetQuotes(w, r)
	case http.MethodPost:
		return s.handleCreateQuote(w, r)
	case http.MethodPut:
		return s.handleUpdateQuote(w, r)
	case http.MethodDelete:
		return s.handleDeleteQuote(w, r)
	default:
		return ErrMethodNotAllowed
	}
}

func (s *ApiServer) handleGetQuotes(w http.ResponseWriter, r *http.Request) error {

	quotes, err := s.store.GetQuotes()
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, quotes)

}

func (s *ApiServer) handleGetQuoteById(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	fmt.Println(id)
	return writeJson(w, http.StatusOK, Quote{})
}

func (s *ApiServer) handleCreateQuote(w http.ResponseWriter, r *http.Request) error {
	createQuoteReq := new(CreateQuoteRequest)
	if err := json.NewDecoder(r.Body).Decode(createQuoteReq); err != nil {
		return err
	}

	quote := NewQuote(createQuoteReq.Author, createQuoteReq.Text)
	if err := s.store.CreateQuote(quote); err != nil {
		return err
	}

	return writeJson(w, http.StatusCreated, quote)
}

func (s *ApiServer) handleUpdateQuote(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *ApiServer) handleDeleteQuote(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func writeJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func makeHttpHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			if apiErr, ok := err.(apiError); ok {
				writeJson(w, apiErr.Status, apiErr)
				return
			}
			writeJson(w, http.StatusInternalServerError, apiError{Err: "internal server error", Status: http.StatusInternalServerError})
		}
	}
}
