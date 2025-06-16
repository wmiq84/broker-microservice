package main

import (
	"fmt"
	"log"
	"net/http"
)

// port to listen on
// in context of Docker, port that http binds to which is container's port
const webPort = "80"

type Config struct{}

// print web port, set servers addr as web port
func main() {
	// of type config
	app := Config{}

	log.Printf("Port %s\n", webPort)

	// define http server
	srv := &http.Server{
		// alternatively, Addr: ":8080"
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server
	// add log.Fatal() optionally
	err := srv.ListenAndServe()
	if err != nil {
		// stops execution
		log.Panic(err)
	}
}
