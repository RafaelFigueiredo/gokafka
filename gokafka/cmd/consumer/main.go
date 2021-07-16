package main

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	log.Println("starting consumer")
	consumer := NewKafkaConsumer()

	topics := []string{"test"}
	err := consumer.SubscribeTopics(topics, nil)

	if err != nil {
		panic(err)
	}
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			panic(err)
		}
		log.Println(string(msg.Value), msg.TopicPartition)
	}
}

func NewKafkaConsumer() *kafka.Consumer {

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "gokafka_kafka_1:9092",
		"client.id":         "goapp-consumer",
		"group.id":          "goapp-group",
	}

	c, err := kafka.NewConsumer(configMap)

	if err != nil {
		log.Println("error consumer", err.Error())
	}
	return c
}
