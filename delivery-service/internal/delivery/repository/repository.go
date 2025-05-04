package repository

import (
	"context"
	"time"

	"delivery-service/internal/delivery/model"
)

// DeliveryRepository — интерфейс для работы с доставками в БД
type DeliveryRepository interface {
	Create(ctx context.Context, d *model.Delivery) error
	GetById(ctx context.Context, id int) (*model.Delivery, error)
	UpdateStatus(ctx context.Context, id int, status model.DeliveryStatus) error
	AssignCourier(ctx context.Context, id int, courierID int) error
	MarkAsDelivered(ctx context.Context, id int, deliveredAt time.Time) error
}
