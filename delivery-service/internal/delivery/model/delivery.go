package model

import "time"

type DeliveryStatus string

const (
	StatusPending   DeliveryStatus = "pending"
	StatusAssigned  DeliveryStatus = "assigned"
	StatusInTransit DeliveryStatus = "in_transit"
	StatusDelivered DeliveryStatus = "delivered"
	StatusFailed    DeliveryStatus = "failed"
)

type DeliveryPriority string

const (
	PriorityNormal  DeliveryPriority = "normal"
	PriorityExpress DeliveryPriority = "express"
)

type Delivery struct {
	ID                    int              `json:"id"`
	OrderID               int              `json:"order_id"`
	CustomerID            int              `json:"customer_id"`
	CourierID             *int             `json:"courier_id,omitempty"` // может быть nil
	Status                DeliveryStatus   `json:"status"`
	Priority              DeliveryPriority `json:"priority"`
	DeliveryAddress       string           `json:"delivery_address"`
	EstimatedDeliveryTime time.Time        `json:"estimated_delivery_time"`
	DeliveredAt           *time.Time       `json:"delivered_at,omitempty"`
	CreatedAt             time.Time        `json:"created_at"`
	UpdatedAt             time.Time        `json:"updated_at"`
}

// OrderCreatedPayload используется в delivery-service для обработки Kafka-события "order-created"
type OrderCreatedPayload struct {
	OrderID    int
	CustomerID int
	Address    string
	Priority   string
}
