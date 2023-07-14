package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pressly/goose/v3"
	"github.com/segmentio/kafka-go"
	"go-kafka/internal/pkg/db"
	"go-kafka/internal/pkg/repository"
	"go-kafka/internal/pkg/repository/postgresql"
	"log"
	"time"
)

const (
	topic     = "order"
	partition = 0
)

func main() {
	ctx := context.Background()

	sqlDb, database, err := db.NewDB(ctx)
	defer database.GetPool(ctx).Close()

	err = goose.Up(sqlDb.DB, "./internal/pkg/db/migrations")

	if err != nil {
		log.Fatalf("невозможно накатить миграции: %v", err)
	}

	ordersRepo := postgresql.NewOrdersRepo(database)

	// to consume messages
	conn, err := kafka.DialLeader(ctx, "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(1e3, 2e3)

	b := make([]byte, 1e3)

	for {
		n, err := batch.Read(b)
		if err != nil {
			fmt.Println(err)
			break
		}

		var order *repository.Order

		err = json.Unmarshal(b[:n], &order)

		if err != nil {
			log.Fatal(err)
		}

		_, err = ordersRepo.Add(ctx, order)

		if err != nil {
			return
		}

	}

	if err = batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err = conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}

	orders, err := ordersRepo.List(ctx)

	for i := 0; i < len(orders); i++ {
		order, _ := json.Marshal(orders[i])

		fmt.Println(string(order))
	}

}
