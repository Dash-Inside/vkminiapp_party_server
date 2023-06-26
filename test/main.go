package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

// a simple handler that sleeps for a duration of time
func timedHandler(duration time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		time.Sleep(duration)
		c.JSON(200, gin.H{"hello": "world"})
	}
}

func main() {

	// create new gin without any middleware
	engine := gin.New()

	// create a route that will last 5 seconds
	engine.GET("/", timedHandler(time.Second*5))

	// create the server
	server := http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	// create an err channel to watch for server errors
	errChan := make(chan error)
	go func(errChan chan error) {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}(errChan)

	// create a quit channel to watch for SIGINT
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// non-blocking select to watch for errors or quit signal
	select {

	// real error logs and exits immediately
	case err := <-errChan:
		log.Fatalln(err)

	// sigint shuts the server down gracefully
	case <-quit:

		// log that we are starting the shut down
		log.Println("Shutting down server after requests finish...")

		// add a backup timeout to the context so if requests
		// don't finish in time they are cut short
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		defer cancel()

		// tell the server to shut down with the backup timeout context
		if err := server.Shutdown(ctx); err != nil {

			// if there's an error just shut down immediately
			log.Fatalln(err)
		}

		// log that the server is shutting down now
		// hopefully with all the requests finished
		log.Println("Server exiting now")
	}
}
