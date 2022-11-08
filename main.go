package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte(`hello world`))
	})

	s := http.Server{Addr: "localhost:8081"}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			// exclude http.ErrServerClosed because it's returned when
			// Shutdown is called.
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println("http server error:", err)
			}
		}
	}()

	// create a context to handle system signal from host machine
	ctx := context.Background()
	ctx, stopFn := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stopFn()

	select {
	case <-ctx.Done():
		log.Println("prepare to shutdown server")
		// wait 15 seconds for inflight requests to complete
		shutdownCtx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFn()
		if err := s.Shutdown(shutdownCtx); err != nil {
			log.Fatal("server shutdown failed:", err)
		}
		log.Println("server shutdown completed.")
	}
}
