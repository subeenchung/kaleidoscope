package main

import (
	"fmt"

	"gitlab.com/jesteringjester/kaleidoscope/types"

	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "testgrp",
		"auto.offset.reset": "latest",
	})
	fmt.Println("consumer established")
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"winlogbeat"}, nil)

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			wl := &types.Winlogbeat{}
			je := json.Unmarshal(msg.Value, &wl)
			fmt.Printf("event id: %v\n", wl.Event.Code)

			//fmt.Printf("Message on %s: %s \n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer err: %v (%v)\n", err, msg)
		}
	}

	c.Close()

}
