package handler

import (
	"context"
	deliveryv1 "delivery-service/api/delivery/v1"
	"delivery-service/internal/delivery/model"
	"delivery-service/internal/delivery/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCHandler struct {
	deliveryv1.UnimplementedDeliveryServiceServer
	service *service.Service
}

func NewGRPCHandler(s *service.Service) *GRPCHandler {
	return &GRPCHandler{service: s}
}

func (h *GRPCHandler) GetDelivery(ctx context.Context, req *deliveryv1.GetDeliveryRequest) (*deliveryv1.GetDeliveryResponse, error) {
	d, err := h.service.GetById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &deliveryv1.GetDeliveryResponse{
		Delivery: convertToProto(d),
	}, nil
}

func (h *GRPCHandler) UpdateStatus(ctx context.Context, req *deliveryv1.UpdateStatusRequest) (*deliveryv1.Empty, error) {
	err := h.service.UpdateStatus(ctx, req.Id, model.DeliveryStatus(req.Status))
	return &deliveryv1.Empty{}, err
}

func (h *GRPCHandler) AssignCourier(ctx context.Context, req *deliveryv1.AssignCourierRequest) (*deliveryv1.Empty, error) {
	err := h.service.AssignCourier(ctx, req.Id, int(req.CourierId))
	return &deliveryv1.Empty{}, err
}

func (h *GRPCHandler) MarkAsDelivered(ctx context.Context, req *deliveryv1.MarkAsDeliveredRequest) (*deliveryv1.Empty, error) {
	err := h.service.MarkAsDelivered(ctx, req.Id, req.DeliveredAt.AsTime())
	return &deliveryv1.Empty{}, err
}

// convertToProto преобразует модель в protobuf-структуру
func convertToProto(d *model.Delivery) *deliveryv1.Delivery {
	var deliveredAt *timestamppb.Timestamp
	if d.DeliveredAt != nil {
		deliveredAt = timestamppb.New(*d.DeliveredAt)
	}

	var courierID int32
	if d.CourierID != nil {
		courierID = int32(*d.CourierID)
	}

	return &deliveryv1.Delivery{
		Id:                    int32(d.ID),
		OrderId:               int32(d.OrderID),
		CustomerId:            int32(d.CustomerID),
		CourierId:             courierID,
		Status:                string(d.Status),
		Priority:              string(d.Priority),
		DeliveryAddress:       d.DeliveryAddress,
		EstimatedDeliveryTime: timestamppb.New(d.EstimatedDeliveryTime),
		DeliveredAt:           deliveredAt,
		CreatedAt:             timestamppb.New(d.CreatedAt),
		UpdatedAt:             timestamppb.New(d.UpdatedAt),
	}
}
