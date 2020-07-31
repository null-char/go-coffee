package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/null-char/go-coffee/handlers"
)

func main() {
	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(logger)
	pHandler := handlers.NewProducts(logger)
	// Create a custom ServeMux for handling requests
	serveMux := http.NewServeMux()
	serveMux.Handle("/", hh)
	serveMux.Handle("/products", pHandler)

	server := &http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	setupCloseHandler(server, logger)
}

func setupCloseHandler(server *http.Server, logger *log.Logger) {
	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt)
	signal.Notify(sc, os.Kill)

	// Block for a signal.
	<-sc
	logger.Println("Server is gracefully shutting down.")
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := server.Shutdown(tc); err != nil {
		logger.Printf("Shut down error: %v", err)
	}
}
