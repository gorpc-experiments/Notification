package main

import (
	"context"
	"github.com/gorpc-experiments/ServiceCore"
	"github.com/redis/go-redis/v9"
	"log"
)

var ctx = context.Background()

func main() {
	ServiceCore.SetupLogging()

	log.Println("Starting up")

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	subscriber := rdb.Subscribe(ctx, "send-user-data")

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(msg.Payload)
	}
}