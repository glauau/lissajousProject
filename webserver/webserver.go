// webserver.go
package main

import (
	"net/http"

	"./lissajous"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	lissajous.Generate(w) // Chame a função lissajous.Generate para gerar a imagem GIF
}
