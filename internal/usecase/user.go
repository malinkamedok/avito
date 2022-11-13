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
	exists, err := u.repo.CheckUserExistence(ctx, user)
	if err != nil {
		return err

	}
	if !exists {
		return u.repo.CreateNewBalance(ctx, user, sum)
	}
	return u.repo.AppendBalance(ctx, user, sum)
}

func (u *UserUseCase) GetBalance(ctx context.Context, uuid uuid.UUID) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) ReserveMoney(ctx context.Context, uuid uuid.UUID, uuid2 uuid.UUID, i int64) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserUseCase) AcceptIncome(ctx context.Context, uuid uuid.UUID, uuid2 uuid.UUID, i int64) error {
	//TODO implement me
	panic("implement me")
}
