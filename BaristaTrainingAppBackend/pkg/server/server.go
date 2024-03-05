package server

import (
	"context"
	"net/http"
	"time"

	"github.com/rs/cors"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost", "http://localhost:3000", "http://localhost:8000"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})

	newHandler := c.Handler(handler)

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        newHandler,
		MaxHeaderBytes: 1 << 20, //1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
