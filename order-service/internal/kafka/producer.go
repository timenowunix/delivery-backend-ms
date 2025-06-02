package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(brokers []string, topic string) *Producer {
	return &Producer{
		writer: &kafka.Writer{
			Addr:         kafka.TCP(brokers...),
			Topic:        topic,
			Balancer:     &kafka.LeastBytes{},
			RequiredAcks: kafka.RequireAll,
		},
	}
}

func (p *Producer) SendOrderCreated(ctx context.Context, event OrderCreatedEvent) error {
	message, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(time.Now().Format(time.RFC3339)),
		Value: message,
	})
	if err != nil {
		log.Printf("Ошибка отправки сообщения: %v", err)
		return err
	}

	log.Printf("OrderCreatedEvent отправлен: %+v", event)
	return nil
}

func (p *Producer) Close() error {
	return p.writer.Close()
}
