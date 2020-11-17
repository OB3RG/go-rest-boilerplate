package auth

import (
	"net/http"

	"github.com/go-chi/chi"

	"go-rest-boilerplate/pkg/config"
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
	a := StaticAuthService()
	h := NewHandler(a)

	s.r.Post("/login", h.handleLogin)

	return s.r
}
