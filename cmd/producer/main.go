package main

import (
	"fmt"
	k "producer/internal/kafka"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const (
	topic        = "my-topic"
	numberOfKeys = 20
)

var address = []string{"localhost:9091", "localhost:9092", "localhost:9093"}

func main() {
	p, err := k.NewProducer(address)
	if err != nil {
		logrus.Fatal(err)
	}

	keys := generateUUIDString()
	for i := 0; i < 100; i++ {
		msg := fmt.Sprintf("Kafka message %d", i)
		key := keys[i%numberOfKeys]
		if err := p.Produce(msg, topic, key); err != nil {
			logrus.Error(err)
		}
	}
}

func generateUUIDString() [numberOfKeys]string {
	var uuids [numberOfKeys]string
	for i := 0; i < numberOfKeys; i++ {
		uuids[i] = uuid.NewString()
	}
	return uuids
}
