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
		time.Sleep(3 * time.Second)
		w.Write([]byte(`hello world`))
	})

	s := http.Server{Addr: "localhost:8081"}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println("http server error:", err)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-c:
		log.Println("prepare to shutdown server")
		ctx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFn()
		if err := s.Shutdown(ctx); err != nil {
			log.Fatal("server shutdown failed:", err)
		}
		log.Println("server shutdown completed.")
	}
}
