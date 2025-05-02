package handler

import (
	"context"
	orderv1 "order-service/api/order/v1"
	"order-service/internal/order/model"
	"order-service/internal/order/service"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderHandler struct {
	orderv1.UnimplementedOrderServiceServer
	orderService *service.OrderService
}

// Конструктор
func NewOrderHandler(orderService *service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

// Реализация методов интерфейса:
func (h *OrderHandler) CreateOrder(ctx context.Context, req *orderv1.CreateOrderRequest) (*orderv1.CreateOrderResponse, error) {
	// Создаем объект модели заказа
	order := &model.Order{
		ParcelID:        req.ParcelId,
		DeliveryAddress: req.DeliveryAddress,
		Status:          req.Status,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	// Вызываем сервис для сохранения заказа
	err := h.orderService.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	// Возвращаем ответ с ID созданного заказа
	return &orderv1.CreateOrderResponse{
		Id: order.ID,
	}, nil
}

func (h *OrderHandler) GetOrder(ctx context.Context, req *orderv1.GetOrderRequest) (*orderv1.GetOrderResponse, error) {
	// Вытаскиваем заказ через сервис
	order, err := h.orderService.GetOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &orderv1.GetOrderResponse{
		Id:              order.ID,
		ParcelId:        order.ParcelID,
		DeliveryAddress: order.DeliveryAddress,
		Status:          order.Status,
		CreatedAt:       timestamppb.New(order.CreatedAt),
		UpdatedAt:       timestamppb.New(order.UpdatedAt),
	}, nil
}

func (h *OrderHandler) UpdateOrderStatus(ctx context.Context, req *orderv1.UpdateOrderStatusRequest) (*orderv1.UpdateOrderStatusResponse, error) {
	// Вызываем сервис для обновления статуса
	err := h.orderService.UpdateOrderStatus(ctx, req.Id, req.Status)
	if err != nil {
		return nil, err
	}

	return &orderv1.UpdateOrderStatusResponse{
		Id: req.Id,
	}, nil
}

func (h *OrderHandler) DeleteOrder(ctx context.Context, req *orderv1.DeleteOrderRequest) (*orderv1.DeleteOrderResponse, error) {
	// Вызываем сервис для удаления заказа
	err := h.orderService.DeleteOrder(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &orderv1.DeleteOrderResponse{
		Id: req.Id,
	}, nil
}
