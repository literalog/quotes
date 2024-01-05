package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/literalog/quotes/pkg/database/mongodb"
	"github.com/literalog/quotes/pkg/quote"
)

type Server struct {
	port     string
	logLevel int
	router   *mux.Router
}

func NewServer(port string) Server {
	s := Server{
		port:     port,
		logLevel: 1,
		router:   mux.NewRouter(),
	}

	storage, err := mongodb.NewMongoStorage()
	if err != nil {
		log.Fatal(err)
	}

	db := storage.Client.Database("quotes")

	quotesRepository := mongodb.NewQuoteRepository(db.Collection("quotes"))
	quotesService := quote.NewService(quotesRepository)
	quotesHandler := quote.NewHandler(quotesService)

	s.router.PathPrefix("/quotes").Handler(quotesHandler.Routes())

	return s
}

func (s *Server) ServeHttp() error {
	log.Println("Server listening on", s.port)
	return http.ListenAndServe(s.port, s.router)
}
