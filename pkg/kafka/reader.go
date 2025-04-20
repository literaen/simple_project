package kafka

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type KFR struct {
	reader *kafka.Reader
}

func NewKafkaReader(brokers []string, topic string, groupID string) *KFR {
	if len(brokers) == 0 {
		panic("Kafka brokers not configured")
	}

	return &KFR{reader: kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic:   topic,
		GroupID: groupID,
	})}
}

func (c *KFR) ReadMessage(ctx context.Context, event interface{}) (string, error) {
	m, err := c.reader.ReadMessage(ctx)
	if err != nil {
		return "", err
	}

	if err := json.Unmarshal(m.Value, &event); err != nil {
		return "", err
	}

	eventType := string(m.Key)

	return eventType, nil
}
