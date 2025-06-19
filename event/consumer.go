package event

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	ampq "github.com/rabbitmq/amqp091-go"
)

type Consumer struct {
	conn      *ampq.Connection
	queueName string
}

func NewConsumer(conn *ampq.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn, // open connection
	}

	// opens new channel, declares exchange
	err := consumer.setup()
	if err != nil {
		return Consumer{}, err
	}

	return consumer, nil
}

func (consumer *Consumer) setup() error {
	// using conn of conn
	// creates channel
	channel, err := consumer.conn.Channel()
	if err != nil {
		return err
	}

	// return result of declaring exchange
	return declareExchange(channel)
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func (consumer *Consumer) Listen(topics []string) error {
	// open new channel, close on ending function
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	// declare random queue for this connection, delete on connection close
	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	// bind queue to logs_topic exchange for each routing key
	// AKA take every message for this exchange (logs_topic), with pattern s and put onto queue
	for _, s := range topics {
		ch.QueueBind(
			q.Name,       // randomly generated
			s,            // iteration through range
			"logs_topic", // name
			false,        // wait
			nil,          // arguments
		)

		if err != nil {
			return err
		}
	}

	// consume message from that queue
	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	// run forever
	// second param left empty, so it's trying to send values to empty
	forever := make(chan bool)
	// go makes it run in background
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)

			go handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message [Exchange Queue] [logs_topic, %s]", q.Name)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	switch payload.Name {
	case "log", "event":
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}

	case "auth":

	default:
		err := logEvent(payload)
		if err != nil {
			log.Println(err)
		}
	}
}

func logEvent(entry Payload) error {
	jsonData, _ := json.MarshalIndent(entry, "", "\t")

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		return err
	}

	// var payload jsonResponse
	// payload.Error = false
	// payload.Message = "logged"

	// app.writeJSON(w, http.StatusAccepted, payload)
	return nil
}
