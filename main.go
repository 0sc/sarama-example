package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/sarama"
)

const topic = "sample-topic"

func main() {
	producer, err := newProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	subscribe(topic, consumer)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello Sarama!") })

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		r.ParseForm()
		msg := prepareMessage(topic, r.FormValue("q"))
		partition, offset, err := producer.SendMessage(msg)
		fmt.Fprintf(w, "Message was saved to partion: %d.\nMessage offset is: %d.\n %s error occured.", partition, offset, err.Error())
	})

	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, html.EscapeString(getMessage())) })

	log.Fatal(http.ListenAndServe(":8081", nil))
}
