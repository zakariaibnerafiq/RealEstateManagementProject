package api

import (
	"database/sql"
	"learninggo/service/image"
	"learninggo/service/user"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
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
	log.Println("Starting server on", s.addr)

	imageHandler := image.NewHandler()
	imageHandler.RegisterRoutes(subrouter)

	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"http://192.168.0.104:5173", "http://localhost:5173"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "UPDATE"}),
	)

	return http.ListenAndServe(s.addr, corsHandler(router))
}
