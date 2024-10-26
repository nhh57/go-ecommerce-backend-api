package initialize

import (
	"github.com/nhh57/go-ecommerce-backend-api/global"
	"github.com/segmentio/kafka-go"
	"log"
)

// Init kafka producer

var KafkaProducer *kafka.Writer

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "otp-auth-topic",
		Balancer: &kafka.LeastBytes{},
	}
}

func CLoseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		log.Fatal("Failed to close kafka producer: %v", err)

	}
}
