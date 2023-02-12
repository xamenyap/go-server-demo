package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(`hello world`))
	})

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Println("http server error:", err)
	}
}