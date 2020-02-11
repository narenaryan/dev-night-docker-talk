package main

import (
	"fmt"
	"log"
	"net/http"
)

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")

}

func main() {
	http.HandleFunc("/health", health)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
