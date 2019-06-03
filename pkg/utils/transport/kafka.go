package transport

import (
	"context"
	"encoding/json"
	"os"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

// BaseKafkaTransport represents a base kafka transport object.
type BaseKafkaTransport struct{}

type handler func(result interface{})

func (t BaseKafkaTransport) Listen(topic string, handler handler) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST")},
		GroupID: os.Getenv("KAFKA_GROUP_ID"),
		Topic:   topic,
	})

	for {
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			break
		}

		var data map[string]interface{}
		err = json.Unmarshal(message.Value, &data)

		if err == nil {
			handler(data)
		}
	}

	reader.Close()
}

func (t BaseKafkaTransport) Send(topic string, payload interface{}) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})

	key, _ := uuid.NewRandom()
	value, _ := json.Marshal(payload)
	writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(key.String()),
		Value: []byte(value),
	})

	writer.Close()
}
