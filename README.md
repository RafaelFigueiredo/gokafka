# gokafka

Kafka project using golang, based on code.education fullcycle course


## How to run?
Up the stack
```bash
docker-compose up --build
```

Attach to kafka container and create a topic
```bash
docker exec -it gokafka_kafka_1 bash
kafka-topics --create --topic test --partitions=2 --bootstrap-server localhost:9092
```

Start consumer
```bash
docker exec -it gokafka_app_1 go run /go/src/cmd/consumer/main.go
```


Start producer
```bash
docker exec -it gokafka_app_1 go run /go/src/cmd/producer/main.go
```

