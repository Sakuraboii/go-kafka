package repository

import (
	"context"
	"errors"
)

var (
	ErrObjectNotFound = errors.New("object not found")
)

type OrderRepo interface {
	Add(ctx context.Context, user *Order) (int64, error)
	List(ctx context.Context) ([]*Order, error)
}
