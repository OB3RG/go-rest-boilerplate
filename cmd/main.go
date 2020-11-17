package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexflint/go-arg"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"go-rest-boilerplate/pkg/config"
	"go-rest-boilerplate/pkg/services/auth"
	"go-rest-boilerplate/pkg/services/user"
)

func main() {
	log.Print("Setting up config")
	cfg := config.DefaultCfg()
	arg.MustParse(cfg)

	log.Print("Setting up new router")
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	userServer := user.NewServer(cfg)
	authServer := auth.NewServer(cfg)

	userRoutes := userServer.SetupRoutes()
	authRoutes := authServer.SetupRoutes()

	r.Mount("/user", userRoutes)
	r.Mount("/auth", authRoutes)

	log.Print("Succesfully loaded routes")

	http.ListenAndServe(":3333", r)
}
