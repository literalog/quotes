package main

import (
	"encoding/json"
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
	r := mux.NewRouter()

	r.HandleFunc("/quotes", makeHttpHandler(s.handleQuotes))
	r.HandleFunc("/quotes/{id}", makeHttpHandler(s.handleQuotes)).Methods(http.MethodGet)

	log.Println("api server listening on", s.listenAddr)

	if err := http.ListenAndServe(s.listenAddr, r); err != nil {
		log.Fatal("error starting api server", err)
	}
}

func (s *ApiServer) handleQuotes(w http.ResponseWriter, r *http.Request) error {
	if id, ok := mux.Vars(r)["id"]; ok {
		return s.handleGetQuoteById(w, r, id)
	}

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
	qq, err := s.store.GetQuotes()
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, qq)
}

func (s *ApiServer) handleGetQuoteById(w http.ResponseWriter, r *http.Request, id string) error {
	q, err := s.store.GetQuoteById(id)
	if err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, q)
}

func (s *ApiServer) handleCreateQuote(w http.ResponseWriter, r *http.Request) error {
	createQuoteReq := new(CreateQuoteRequest)
	if err := json.NewDecoder(r.Body).Decode(createQuoteReq); err != nil {
		return err
	}

	q := NewQuote(createQuoteReq.Author, createQuoteReq.Text)
	if err := s.store.CreateQuote(q); err != nil {
		return err
	}

	return writeJson(w, http.StatusCreated, q)
}

func (s *ApiServer) handleUpdateQuote(w http.ResponseWriter, r *http.Request) error {
	updateQuoteReq := new(UpdateQuoteRequest)
	if err := json.NewDecoder(r.Body).Decode(updateQuoteReq); err != nil {
		return err
	}

	q := NewQuote(updateQuoteReq.Author, updateQuoteReq.Text)
	if err := s.store.UpdateQuote(q); err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, q)
}

func (s *ApiServer) handleDeleteQuote(w http.ResponseWriter, r *http.Request) error {
	deleteQuoteReq := new(DeleteQuoteRequest)
	if err := json.NewDecoder(r.Body).Decode(deleteQuoteReq); err != nil {
		return err
	}

	if err := s.store.DeleteQuote(deleteQuoteReq.Id); err != nil {
		return err
	}

	return writeJson(w, http.StatusOK, nil)
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
