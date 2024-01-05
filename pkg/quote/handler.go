package quote

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/literalog/cerrors"
)

type Handler interface {
	Routes() *mux.Router
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
	router  *mux.Router
}

func NewHandler(s Service) Handler {
	h := &handler{
		service: s,
		router:  mux.NewRouter(),
	}

	h.setupRoutes()

	return h
}

func (h *handler) setupRoutes() {
	h.router.HandleFunc("/", h.Create).Methods(http.MethodPost)
	h.router.HandleFunc("/{id}", h.Update).Methods(http.MethodPut)
	h.router.HandleFunc("/{id}", h.Delete).Methods(http.MethodDelete)
	h.router.HandleFunc("/", h.GetAll).Methods(http.MethodGet)
	h.router.HandleFunc("/{id}", h.GetById).Methods(http.MethodGet)
}

func (h *handler) Routes() *mux.Router {
	return h.router
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(QuoteRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		cerrors.Handle(err, w)
		return
	}

	q := NewQuote(*req)
	if err := h.service.Create(ctx, q); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(q)
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	req := new(QuoteRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		cerrors.Handle(err, w)
		return
	}

	q := NewQuote(*req)
	if err := h.service.Update(ctx, q); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(req)
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	if err := h.service.Delete(ctx, id); err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("quote deleted successfully")
}

func (h *handler) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	qq, err := h.service.GetAll(ctx)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(qq)
}

func (h *handler) GetById(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := mux.Vars(r)["id"]

	q, err := h.service.GetById(ctx, id)
	if err != nil {
		cerrors.Handle(err, w)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(q)
}
