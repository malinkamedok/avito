package usecase

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type UserUseCase struct {
	repo UserRp
}

var _ UserContract = (*UserUseCase)(nil)

func NewUserUseCase(repo UserRp) *UserUseCase {
	return &UserUseCase{
		repo: repo,
	}
}

func (u *UserUseCase) AppendBalance(ctx context.Context, user uuid.UUID, sum uint64) error {
	if sum <= 0 {
		log.Println("Cannot work with this sum number")
		return fmt.Errorf("error in sum value")
	}
	exists, err := u.repo.CheckUserBalanceExistence(ctx, user)
	if err != nil {
		return err
	}
	if !exists {
		return u.repo.CreateNewBalance(ctx, user, sum)
	}
	return u.repo.AppendBalance(ctx, user, sum)
}

func (u *UserUseCase) GetBalance(ctx context.Context, user uuid.UUID) (int64, error) {
	exists, err := u.repo.CheckUserBalanceExistence(ctx, user)
	if err != nil {
		return -1, err
	}
	if !exists {
		var sum uint64
		err = u.repo.CreateNewBalance(ctx, user, sum)
		if err != nil {
			return -1, err
		}
	}
	return u.repo.GetBalance(ctx, user)
}

func (u *UserUseCase) GetReserve(ctx context.Context, user uuid.UUID) (int64, error) {
	exists, err := u.repo.CheckUserReserveExistence(ctx, user)
	if err != nil {
		return -1, err
	}
	if !exists {
		fmt.Errorf("user does not have any reserved maney %w", err)
	}
	return u.repo.GetReserve(ctx, user)
}

func (u *UserUseCase) ReserveMoney(ctx context.Context, balanceUUID uuid.UUID, reserveUUID uuid.UUID, amount uint64) error {
	exists, err := u.repo.CheckUserReserveExistence(ctx, reserveUUID)
	if err != nil {
		return err
	}
	if !exists {
		return u.repo.CreateNewReserve(ctx, balanceUUID, amount)
	}
	enough, err := u.repo.CheckEnoughMoneyBalance(ctx, balanceUUID, amount)
	if err != nil {
		return err
	}
	if !enough {
		return fmt.Errorf("not enough money %w", err)
	}
	return u.repo.ReserveMoney(ctx, balanceUUID, reserveUUID, amount)
}

func (u *UserUseCase) AcceptIncome(ctx context.Context, uuid uuid.UUID, uuid2 uuid.UUID, i int64) error {
	//TODO implement me
	panic("implement me")
}
