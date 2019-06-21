package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Jean Saunie's Website\n")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello world!\n")

		if err != nil {
			println("Error")
		}
	})

	log.Fatal(http.ListenAndServe(":8001", nil))
}
