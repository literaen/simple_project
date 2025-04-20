package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type KFW struct {
	writer *kafka.Writer
}

func NewKafkaWriter(brokers []string, topic string) *KFW {
	if len(brokers) == 0 {
		panic("Kafka brokers not configured")
	}

	return &KFW{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: brokers,
			Topic:   topic,
		}),
	}
}

func (k *KFW) WriteString(ctx context.Context, key []byte, value string) error {
	msg := kafka.Message{
		Key:   key,
		Value: []byte(value),
	}

	err := k.writer.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}

func (k *KFW) Write(ctx context.Context, key []byte, value interface{}) error {
	message, err := json.Marshal(value)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   key,
		Value: message,
	}

	err = k.writer.WriteMessages(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
