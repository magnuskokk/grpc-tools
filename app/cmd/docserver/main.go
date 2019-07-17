package main

import (
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
)

func main() {
	swagger := packr.NewBox(os.Getenv("GOPATH") + "/swagger")
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(swagger)))
	http.ListenAndServe(":8000", nil)
}
