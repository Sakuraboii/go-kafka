package repository

const (
	OrderStateCreate   = "create"   // Заказ создается
	OrderStateProgress = "progress" // Заказ выполняется
	OrderStateComplete = "complete" // Заказ выполнен
)

type Order struct {
	Id           int64  `db:"id"`
	ItemId       int64  `db:"item_id"`
	UserId       int64  `db:"user_id"`
	OrderPointId int64  `db:"orderpoint_id"`
	OrderState   string `db:"orderstate"`
}
