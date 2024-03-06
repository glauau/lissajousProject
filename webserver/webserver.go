package main

import (
	"fmt"
	"net/http"
)

func OutGif() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "out.gif")
	})
	fmt.Println("Servindo em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	OutGif()
}
