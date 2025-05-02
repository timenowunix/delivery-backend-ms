package model

import "time"

type Order struct {
	ID              int64     `db:"id"`
	ParcelID        int64     `db:"parcel_id"`
	DeliveryAddress string    `db:"delivery_address"`
	Status          string    `db:"status"`
	CreatedAt       time.Time `db:"created_at"`
	UpdatedAt       time.Time `db:"updated_at"`
}
