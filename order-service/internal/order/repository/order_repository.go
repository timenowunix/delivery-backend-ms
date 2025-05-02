package repository

import (
	"context"
	"order-service/internal/order/model"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	return &OrderRepository{db: db}
}

// CreateOrder сохраняет новый заказ в базе
func (r *OrderRepository) CreateOrder(ctx context.Context, order *model.Order) error {
	query := `
		INSERT INTO orders (parcel_id, delivery_address, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.db.QueryRow(ctx, query,
		order.ParcelID,
		order.DeliveryAddress,
		order.Status,
		order.CreatedAt,
		order.UpdatedAt,
	).Scan(&order.ID)

	return err
}

// GetOrder получает заказ по ID
func (r *OrderRepository) GetOrder(ctx context.Context, id int64) (*model.Order, error) {
	query := `
	SELECT id, parcel_id, delivery_address, status, created_at, updated_at
	FROM orders
	WHERE id = $1
	`

	order := &model.Order{}
	err := r.db.QueryRow(ctx, query, id).Scan(
		&order.ID,
		&order.ParcelID,
		&order.DeliveryAddress,
		&order.Status,
		&order.CreatedAt,
		&order.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return order, nil
}

// UpdateOrderStatus обновляет статус заказа
func (r *OrderRepository) UpdateOrderStatus(ctx context.Context, id int64, status string) error {
	query := `
	UPDATE orders
	SET status = $1, updated_at = $2
	WHERE id = $3
	`

	_, err := r.db.Exec(ctx, query, status, time.Now(), id)
	return err
}

// DeleteOrder удаляет заказ по ID
func (r *OrderRepository) DeleteOrder(ctx context.Context, id int64) error {
	query := `
	DELETE FROM orders
	WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, id)
	return err
}
