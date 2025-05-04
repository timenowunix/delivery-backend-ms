package repository

import (
	"context"
	"fmt"
	"time"

	"delivery-service/internal/delivery/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PGXRepository struct {
	pool *pgxpool.Pool
}

func NewPgxRepository(pool *pgxpool.Pool) *PGXRepository {
	return &PGXRepository{pool: pool}
}

func (r *PGXRepository) Create(ctx context.Context, d *model.Delivery) error {
	query := `
		INSERT INTO deliveries (
			order_id, customer_id, courier_id, status, priority, delivery_address,
			estimated_delivery_time, delivered_at, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		 RETURNING id;
	`

	err := r.pool.QueryRow(ctx, query,
		d.OrderID, d.CustomerID, d.CourierID, d.Status, d.Priority, d.DeliveryAddress, d.EstimatedDeliveryTime, d.DeliveredAt, d.CreatedAt, d.UpdatedAt,
	).Scan(&d.ID)
	if err != nil {
		return fmt.Errorf("failed to create delivery: %w", err)
	}
	return nil
}

func (r *PGXRepository) GetById(ctx context.Context, id int) (*model.Delivery, error) {
	query := `
		SELECT id, order_id, customer_id, courier_id, status, priority,
		       delivery_address, estimated_delivery_time, delivered_at,
		       created_at, updated_at
		FROM deliveries WHERE id = $1;
	`

	var d model.Delivery
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&d.ID, &d.OrderID, &d.CustomerID, &d.CourierID, &d.Status, &d.Priority,
		&d.DeliveryAddress, &d.EstimatedDeliveryTime, &d.DeliveredAt,
		&d.CreatedAt, &d.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get delivery by id: %w", err)
	}
	return &d, nil
}

func (r *PGXRepository) UpdateStatus(ctx context.Context, id int, status model.DeliveryStatus) error {
	query := `UPDATE deliveries SET status = $1, updated_at = $2 WHERE id = $3`
	_, err := r.pool.Exec(ctx, query, status, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to update status: %w", err)
	}
	return nil
}

func (r *PGXRepository) AssignCourier(ctx context.Context, id int, courierID int) error {
	query := `UPDATE deliveries SET courier_id = $1, status = $2, updated_at = $3 WHERE id = $4`
	_, err := r.pool.Exec(ctx, query, courierID, model.StatusAssigned, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to assign courier: %w", err)
	}
	return nil
}

func (r *PGXRepository) MarkAsDelivered(ctx context.Context, id int, deliveredAt time.Time) error {
	query := `UPDATE deliveries SET status = $1, delivered_at = $2, updated_at = $3 WHERE id = $4`
	_, err := r.pool.Exec(ctx, query, model.StatusDelivered, deliveredAt, time.Now(), id)
	if err != nil {
		return fmt.Errorf("failed to mark as delivered: %w", err)
	}
	return nil
}

var _ DeliveryRepository = (*PGXRepository)(nil)
