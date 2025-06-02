package service

import (
	"context"
	"delivery-service/internal/delivery/model"
	"delivery-service/internal/delivery/repository"

	"time"
)

type Service struct {
	repo repository.DeliveryRepository
}

func NewService(repo repository.DeliveryRepository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetById(ctx context.Context, id int32) (*model.Delivery, error) {
	return s.repo.GetById(ctx, int(id))
}

func (s *Service) UpdateStatus(ctx context.Context, id int32, status model.DeliveryStatus) error {
	return s.repo.UpdateStatus(ctx, int(id), status)
}

func (s *Service) AssignCourier(ctx context.Context, id int32, courierID int) error {
	return s.repo.AssignCourier(ctx, int(id), courierID)
}

func (s *Service) MarkAsDelivered(ctx context.Context, id int32, deliveredAt time.Time) error {
	return s.repo.MarkAsDelivered(ctx, int(id), deliveredAt)
}

// CreateFromEvent создает доставку на основе события от order-service (Kafka: "order-created")
func (s *Service) CreateFromEvent(ctx context.Context, e model.OrderCreatedPayload) error {
	delivery := &model.Delivery{
		OrderID:         e.OrderID,
		CustomerID:      e.CustomerID,
		Status:          model.StatusPending,
		Priority:        model.DeliveryPriority(e.Priority),
		DeliveryAddress: e.Address,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	return s.repo.Create(ctx, delivery)
}
