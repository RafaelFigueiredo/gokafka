package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	log.Println("starting")

	producer := NewKafkaProducer()
	defer producer.Flush(10)

	deliveryChan := make(chan kafka.Event)

	go DeliveryReport(deliveryChan)
	i := 0
	for {
		Publish("hello world "+fmt.Sprint(i), "test", producer, nil, deliveryChan)
		time.Sleep(5 * time.Millisecond)
		i++
	}
}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "gokafka_kafka_1:9092",
		"delivery.timeout.ms": "1",
		"acks":                "all", //0-no ack, 1-leader, all,
		"enable.idempotence":  "true",
	}
	p, err := kafka.NewProducer(configMap)

	if err != nil {
		log.Println(err.Error())
	}

	return p
}

func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Key:            key,
	}

	err := producer.Produce(message, deliveryChan)

	if err != nil {
		return err
	}

	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) {
	for e := range deliveryChan {
		switch e.(type) {
		case *kafka.Message:
			e := <-deliveryChan
			msg := e.(*kafka.Message)

			if msg.TopicPartition.Error != nil {
				log.Println("Erro ao enviar")
			} else {
				log.Println("Mensagem enviada", msg.TopicPartition)
				// anotar no banco de dados que a mensagem foi processada.
				// ex: confirma que uma transferencia bancaria ocorreu.
			}
		}
	}
}
