package postgresql

import (
	"context"
	"go-kafka/internal/pkg/db"
	"go-kafka/internal/pkg/repository"
)

type OrderRepo struct {
	db db.DBops
}

func NewOrdersRepo(db db.DBops) *OrderRepo {
	return &OrderRepo{db: db}
}

func (r *OrderRepo) Add(ctx context.Context, order *repository.Order) (int64, error) {
	var id int64
	err := r.db.ExecQueryRow(ctx,
		`INSERT INTO orders(item_id,user_id,orderpoint_id,orderstate) VALUES ($1,$2,$3,$4) RETURNING id`,
		order.ItemId,
		order.UserId,
		order.OrderPointId,
		order.OrderState).Scan(&id)
	return id, err
}

func (r *OrderRepo) List(ctx context.Context) ([]*repository.Order, error) {
	orders := make([]*repository.Order, 0)
	err := r.db.Select(ctx, &orders,
		"SELECT id,item_id,user_id,orderpoint_id,orderstate FROM orders")
	return orders, err
}
