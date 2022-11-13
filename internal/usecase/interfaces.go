package usecase

import (
	"context"
	"github.com/google/uuid"
)

type (
	UserRp interface {
		CheckUserExistence(context.Context, uuid.UUID) (bool, error)
		CreateNewBalance(context.Context, uuid.UUID, uint64) error
		AppendBalance(context.Context, uuid.UUID, uint64) error
		GetBalance(context.Context, uuid.UUID) (int64, error)
		ReserveMoney(context.Context, uuid.UUID, uuid.UUID, int64) error
		AcceptIncome(context.Context, uuid.UUID, uuid.UUID, int64) error
	}

	UserContract interface {
		AppendBalance(context.Context, uuid.UUID, uint64) error
		GetBalance(context.Context, uuid.UUID) (int64, error)
		ReserveMoney(context.Context, uuid.UUID, uuid.UUID, int64) error
		AcceptIncome(context.Context, uuid.UUID, uuid.UUID, int64) error
	}
)
