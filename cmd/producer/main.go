package main

import (
	"context"
	"encoding/json"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"go-kafka/internal/pkg/repository"
	"log"
	"math/rand"
	"time"
)

const (
	topic     = "order"
	partition = 0
)

func main() {
	// to produce messages

	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	err = conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		return
	}

	for i := 0; i < 15; i++ {

		order := repository.Order{
			ItemId:       int64(rand.Intn(50)),
			UserId:       int64(rand.Intn(50)),
			OrderPointId: 443124,
			OrderState:   "Samara",
		}

		data, err := json.Marshal(order)

		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write(data)

		if err != nil {
			return
		}

	}

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
