package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response(w, r)
	})

	log.Println("server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func response(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello from Go on Elastic Beanstalk")
	if err != nil {
		log.Fatal(err)
	}
}
