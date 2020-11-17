package user

import (
	"net/http"

	"github.com/go-chi/chi"

	"go-rest-boilerplate/pkg/config"
	"go-rest-boilerplate/pkg/middleware"
)

// Server ...
type Server struct {
	cfg *config.Config
	r   *chi.Mux
}

// NewServer ...
func NewServer(cfg *config.Config) *Server {
	return &Server{cfg: cfg}
}

// SetupRoutes ...
func (s *Server) SetupRoutes() http.Handler {
	s.r = chi.NewRouter()

	handler := NewHandler()

	s.r.Get("/", middleware.IsAdmin(handler.handleIndex))

	return s.r
}
