package kafka

type OrderCreatedEvent struct {
	OrderID    int    `json:"order_id"`
	CustomerID int    `json:"customer_id"`
	ParcelID   int    `json:"parcel_id"`
	Address    string `json:"delivery_address"`
	Priority   string `json:"priority"` // "normal" или "express"
}
