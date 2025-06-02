package kafka

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"delivery-service/internal/delivery/model"
	"delivery-service/internal/delivery/service"
	"delivery-service/internal/event"

	"github.com/segmentio/kafka-go"
)

func StartOrderConsumer(ctx context.Context, svc *service.Service) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "order-created",
		GroupID:  "delivery-service",
		MinBytes: 10e3,
		MaxBytes: 10e6,
		MaxWait:  1 * time.Second,
	})

	go func() {
		defer r.Close()
		log.Println("Kafka consumer started: topic order-created")

		for {
			m, err := r.ReadMessage(ctx)
			if err != nil {
				log.Printf("Kafka read error: %v", err)
				break
			}

			var orderEvent event.OrderCreatedEvent
			if err := json.Unmarshal(m.Value, &orderEvent); err != nil {
				log.Printf("JSON decode error: %v", err)
				continue
			}

			log.Printf("Received order-created: %+v", orderEvent)

			payload := model.OrderCreatedPayload{
				OrderID:    orderEvent.OrderID,
				CustomerID: orderEvent.CustomerID,
				Address:    orderEvent.Address,
				Priority:   orderEvent.Priority,
			}
			// Вызов сервиса для создания доставки
			err = svc.CreateFromEvent(ctx, payload)
			if err != nil {
				log.Printf("Failed to create delivery: %v", err)
			}
		}
	}()
}
