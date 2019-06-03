package transport

import (
	"context"
	"encoding/json"
	"os"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

// BaseKafkaTransport represents a base kafka transport object.
type BaseKafkaTransport struct{}

type handler func(result interface{})

func (t BaseKafkaTransport) Listen(ctx context.Context, topic string, handler handler) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST")},
		GroupID: os.Getenv("KAFKA_GROUP_ID"),
		Topic:   topic,
	})
	defer reader.Close()

	for {
		message, err := reader.ReadMessage(ctx)
		if err != nil {
			log.WithError(err).Warn("An error occurred while reading the message")
			break
		}

		var data map[string]interface{}
		err = json.Unmarshal(message.Value, &data)
		if err != nil {
			log.WithError(err).Warn("An error occurred while converting the message")
			break
		}

		handler(data)
	}
}

func (t BaseKafkaTransport) Send(ctx context.Context, topic string, payload interface{}) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	key, err := uuid.NewRandom()
	if err != nil {
		log.WithError(err).Warn("An error occurred while generating the key")
		return
	}

	value, err := json.Marshal(payload)
	if err != nil {
		log.WithError(err).Warn("An error occurred while converting the message")
		return
	}

	writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(key.String()),
		Value: []byte(value),
	})
}
