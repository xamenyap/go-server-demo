package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		log.Println("request received by default server")
		time.Sleep(50 * time.Millisecond)
		w.Write([]byte(`hello world`))
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("http server error:", err)
	}
}