package main

import (
	"mime"
	"net/http"

	"github.com/prometheus/common/log"
	"github.com/rakyll/statik/fs"

	// Static files
	_ "grpc-tools/statik"
)

// serveOpenAPI serves an OpenAPI UI on /openapi-ui/
func serveOpenAPI(mux *http.ServeMux) error {
	mime.AddExtensionType(".svg", "image/svg+xml")

	statikFS, err := fs.New()
	if err != nil {
		return err
	}

	// Expose files in static on <host>/openapi-ui
	fileServer := http.FileServer(statikFS)
	prefix := "/openapi-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	return nil
}

func main() {
	mux := http.NewServeMux()
	if err := serveOpenAPI(mux); err != nil {
		log.Fatalln("Failed to serve OpenAPI UI")
	}

	log.Info("Serving OpenAPI Documentation on localhost:8000/openapi-ui/")
	srv := http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
	log.Fatalln(srv.ListenAndServe())
}
