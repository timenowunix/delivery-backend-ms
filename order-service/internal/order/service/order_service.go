package service

import (
	"context"
	"order-service/internal/order/model"
	"order-service/internal/order/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

// CreateOrder создает новый заказ
func (s *OrderService) CreateOrder(ctx context.Context, order *model.Order) error {
	return s.repo.CreateOrder(ctx, order)
}

// GetOrder получает заказ по ID
func (s *OrderService) GetOrder(ctx context.Context, id int64) (*model.Order, error) {
	return s.repo.GetOrder(ctx, id)
}

// UpdateOrderStatus обновляет статус заказа
func (s *OrderService) UpdateOrderStatus(ctx context.Context, id int64, status string) error {
	return s.repo.UpdateOrderStatus(ctx, id, status)
}

// DeleteOrder удаляет заказ по ID
func (s *OrderService) DeleteOrder(ctx context.Context, id int64) error {
	return s.repo.DeleteOrder(ctx, id)
}
