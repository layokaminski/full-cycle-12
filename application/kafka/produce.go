package kafka

import (
	"encoding/json"
	route2 "github.com/layokaminski/imersaofsfc2-simulator/application/route"
	"github.com/layokaminski/imersaofsfc2-simulator/infra/kafka"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	
	json.Unmarshal(msg.Value, &route)

	route.LoadPositions()
	
	positions, err := route.ExportJsonPositions()

	if err != nil {
		log.Println(err.Error())
	}

	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProducerTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}