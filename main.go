package main

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"

	"fmt"
	"net/http"
	// "time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	log.WithFields(log.Fields{}).Info("test")

	fmt.Fprintf(w, "sauce\n")
}
func main() {
	log.SetFormatter(&log.JSONFormatter{})

	http.HandleFunc("/sauce", hello)
	_ = http.ListenAndServe(":80", nil)

	client := redis.NewClient(&redis.Options{
		Addr:       "sage-redis:6379",
		DB:         1,
		MaxRetries: 5,
	})

	pong, err := client.Ping().Result()

	if err != nil {
		panic(err)
	}

	log.Info("redis:", pong)

	pubsub := client.Subscribe("mychannel1")

	// Wait for confirmation that subscription is created before publishing anything.
	_, err = pubsub.Receive()
	if err != nil {
		panic(err)
	}

	// Go channel which receives messages.
	ch := pubsub.Channel()

	// Publish a message.
	err = client.Publish("mychannel1", "hello").Err()
	if err != nil {
		panic(err)
	}

	// time.AfterFunc(time.Second, func() {
	// 	// When pubsub is closed channel is closed too.
	// 	_ = pubsub.Close()
	// })

	// Consume messages.
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
