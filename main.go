package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(3 * time.Second)
		w.Write([]byte(`hello world`))
	})

	s := http.Server{
		Addr: "localhost:8081",
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println("http server error:", err)
			}
		}
	}()

	signal := make(chan os.Signal, 1)

	select {
	case <-signal:
		log.Println("prepare to shutdown server")
		ctx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFn()
		if err := s.Shutdown(ctx); err != nil {
			log.Fatal("server shutdown failed:", err)
		}
		log.Println("server shutdown completed.")
	}
}
