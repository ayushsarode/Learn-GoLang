package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	startDelay time.Duration
}

func NewServer(addr string, startDelay time.Duration) *Server {
	return &Server{
		startDelay: startDelay,
		httpServer: &http.Server{
			Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			}),
		},
	}
}

func(s *Server) Start(ctx context.Context) error {
	select {
	case <- time.After(s.startDelay):
	case <- ctx.Done():
		return errors.New("error startup exceeded time limit")
	}

	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println("server error: %v", err)
		}
	}()

	return nil
}

func main() {
	fmt.Println("Hello")
	

}

func StartServer(ctx context.Context) error {
	

