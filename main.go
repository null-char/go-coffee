package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/null-char/go-coffee/docs"
	"github.com/null-char/go-coffee/products"
)

const local = "127.0.0.1:9090"

func main() {
	logger := log.New(os.Stdout, "[GLOBAL] ", log.LstdFlags)
	r := gin.Default()

	products.RegisterRoutes(r.Group("/products"))
	// Documentation will be served on "/docs"
	docs.RegisterRoutes(r.Group("/"))

	server := &http.Server{
		Addr:         local,
		Handler:      r,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	// Deal with listening in on incoming connections in a different goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal(err)
		}
	}()

	// This function will essentially block the main goroutine and wait until we receive either
	// a SIGINT or SIGKILL and gracefully shutdown the server if so.
	setupCloseHandler(server, logger)
}

func setupCloseHandler(server *http.Server, logger *log.Logger) {
	sc := make(chan os.Signal)
	signal.Notify(sc, os.Interrupt, os.Kill)

	// Block for a signal.
	<-sc
	logger.Println("Server is gracefully shutting down.")
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	if err := server.Shutdown(tc); err != nil {
		logger.Printf("Shut down error: %v", err)
	}
}
