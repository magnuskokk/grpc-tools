package main

import (
	"net/http"
	"os"

	"github.com/gobuffalo/packr"
)

func main() {
	ui := packr.NewBox(os.Getenv("GOPATH") + "/swagger")
	json := packr.NewBox(os.Getenv("REPO_ROOT") + "/swagger")
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(ui)))
	http.Handle("/json/", http.StripPrefix("/json/", http.FileServer(json)))
	http.ListenAndServe(":8000", nil)
}
