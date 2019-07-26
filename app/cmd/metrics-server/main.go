package main

import (
	"app/pkg/metrics"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"log"
	"time"
)

var sigs chan os.Signal
var httpAddr = ":8002"

func init() {
	sigs = make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
}

func main() {
	log.Println("Metrics server listening at :8002")

	srv := &http.Server{
		Addr:         httpAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      metrics.InitHandler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				log.Fatal(err)
			}
		}
	}()

	<-sigs
	log.Println("Shutting down metrics server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Error shutting down http server:", err)
	}
}
