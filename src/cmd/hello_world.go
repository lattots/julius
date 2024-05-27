package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 3000

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/", handleHello)

	fmt.Printf("Server started on port %d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		log.Fatalln(err)
	}
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello from Julius!")
	if err != nil {
		log.Fatalln(err)
	}
}
