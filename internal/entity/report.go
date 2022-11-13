package entity

import "github.com/google/uuid"

type Report struct {
	ID        uint64    `json:"id"`
	UserID    uuid.UUID `json:"user_uuid"`
	ServiceID uuid.UUID `json:"service_uuid"`
	OrderID   uuid.UUID `json:"order_uuid"`
	TotalCost int64     `json:"total_cost"`
}
