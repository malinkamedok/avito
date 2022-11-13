package repo

import (
	"avito/internal/usecase"
	"avito/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
)

type UserRepo struct {
	*postgres.Postgres
}

var _ usecase.UserRp = (*UserRepo)(nil)

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

// add here zerolog

func (u *UserRepo) CheckUserExistence(ctx context.Context, user uuid.UUID) (bool, error) {
	query := `select exists(select * from balance where user_uuid = $1)`

	rows, err := u.Pool.Query(ctx, query, user)
	if err != nil {
		log.Println("Cannot execute query")
		return false, fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()

	var exists bool
	for rows.Next() {
		err = rows.Scan(&exists)
		if err != nil {
			log.Println("cannot scan exists flag")
			return false, fmt.Errorf("cannot scan value")
		}
	}
	return exists, nil
}

func (u *UserRepo) CreateNewBalance(ctx context.Context, user uuid.UUID, sum uint64) error {
	query := `INSERT INTO balance(user_uuid, balance) VALUES($1, $2)`
	rows, err := u.Pool.Query(ctx, query, user, sum)
	if err != nil {
		log.Println("Cannot execute query to create")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	return nil
}

func (u *UserRepo) AppendBalance(ctx context.Context, user uuid.UUID, sum uint64) error {
	query := `SELECT update_balance($1, $2)`

	rows, err := u.Pool.Query(ctx, query, user, sum)
	if err != nil {
		log.Println("Cannot execute query to append")
		return fmt.Errorf("error in executing query %w", err)
	}
	defer rows.Close()
	log.Println("Successfully executed AppendBalance query")
	return nil
}

func (u *UserRepo) GetBalance(ctx context.Context, uuid uuid.UUID) (int64, error) {
	query := `SELECT balance FROM balance WHERE user_uuid = $1`

	rows, err := u.Pool.Query(ctx, query, uuid)
	if err != nil {
		log.Println("Cannot execute query to get user balance")
		return -1, err
	}
	defer rows.Close()
	log.Println("Successfully executed GetBalance query")

	var balance int64
	for rows.Next() {
		err = rows.Scan(&balance)
		if err != nil {
			log.Println("cannot scan balance")
			return -1, fmt.Errorf("cannot scan value")
		}
	}
	return balance, nil
}

func (u *UserRepo) ReserveMoney(ctx context.Context, uuid uuid.UUID, uuid2 uuid.UUID, i int64) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepo) AcceptIncome(ctx context.Context, uuid uuid.UUID, uuid2 uuid.UUID, i int64) error {
	//TODO implement me
	panic("implement me")
}
