package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// port to listen on
// in context of Docker, port that http binds to which is container's port
const webPort = "80"

type Config struct {
	Rabbit *amqp.Connection
}

// print web port, set servers addr as web port
func main() {
	rabbitCon, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitCon.Close()
	// of type config
	app := Config{
		Rabbit: rabbitCon,
	}

	log.Printf("Port %s\n", webPort)

	// define http server
	srv := &http.Server{
		// alternatively, Addr: ":8080"
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// start server
	// add log.Fatal() optionally
	err = srv.ListenAndServe()
	if err != nil {
		// stops execution
		log.Panic(err)
	}
}

func connect() (*amqp.Connection, error) {
	// comverted to int64 b/c of time.Second
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// while true
	for {
		// try to connect, then break (rabbitmq in Docker)
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")
		if err != nil {
			fmt.Println("RabbitMQ not ready yet...")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		// after each failure, break for a bit
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
