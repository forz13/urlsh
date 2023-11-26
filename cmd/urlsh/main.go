package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"urlsh/internal/config"
)

func main() {
	log.Printf("start")
	cfg := config.Load()
	fmt.Println(cfg)
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := &http.Server{
		Addr:         cfg.AppHost + ":" + cfg.AppPort,
		Handler:      r,
		ReadTimeout:  cfg.AppReadTimeout,
		WriteTimeout: cfg.AppWriteTimeout,
		IdleTimeout:  cfg.AppIdleTimeout,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("failed to start server", err)
		}
	}()

	log.Printf("server started")
	<-done
	log.Printf("stopping server")

	// TODO: move timeout to config
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("failed to stop server", err)

		return
	}

	// TODO: close storage

	log.Printf("server stopped")
}
