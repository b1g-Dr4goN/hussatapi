package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/b1g-Dr4goN/hussatapi/internal/auth"
	"github.com/b1g-Dr4goN/hussatapi/internal/user"
	"github.com/b1g-Dr4goN/hussatapi/pkg/middlewares"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	authHandler := auth.NewHandler(userStore)
	authHandler.RegisterRoutes(subrouter)

	handler := middlewares.CorsMiddleware(router)
	handler = middlewares.LoggingMiddleware(handler)

	log.Println("Listening on port", s.addr)

	return http.ListenAndServe(s.addr, handler)
}
