package usecase

import (
	"context"
	"github.com/google/uuid"
)

type (
	UserRp interface {
		CheckUserBalanceExistence(context.Context, uuid.UUID) (bool, error)
		CheckUserReserveExistence(context.Context, uuid.UUID) (bool, error)
		CheckRequiredReserveExistence(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) (bool, error)
		CreateNewBalance(context.Context, uuid.UUID, uint64) error
		CheckEnoughMoneyBalance(context.Context, uuid.UUID, uint64) (bool, error)
		CreateNewReserve(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		AppendBalance(context.Context, uuid.UUID, uint64) error
		GetBalance(context.Context, uuid.UUID) (int64, error)
		GetReserve(context.Context, uuid.UUID) ([]int64, error)
		ReserveMoney(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		AcceptIncome(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
	}

	UserContract interface {
		AppendBalance(context.Context, uuid.UUID, uint64) error
		GetBalance(context.Context, uuid.UUID) (int64, error)
		ReserveMoney(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
		GetReserve(context.Context, uuid.UUID) ([]int64, error)
		AcceptIncome(context.Context, uuid.UUID, uuid.UUID, uuid.UUID, uint64) error
	}
)
