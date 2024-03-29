package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		log.Println("request received by graceful server")
		time.Sleep(200 * time.Millisecond)
		w.Write([]byte(`hello world`))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`ok`))
	})

	s := http.Server{Addr: ":8081"}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			// exclude http.ErrServerClosed because it's returned when
			// Shutdown is called.
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println("http server error:", err)
			}
		}
	}()

	// create a buffered channel to handle
	// system signal from host machine.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-c:
		log.Println("prepare to shutdown server")
		// wait 15 seconds for inflight requests to complete
		ctx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFn()
		if err := s.Shutdown(ctx); err != nil {
			log.Fatal("server shutdown failed:", err)
		}
		log.Println("server shutdown completed.")
	}
}
