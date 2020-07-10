package main

import (
	"fmt"
	"log"

	"encoding/json"
	"io/ioutil"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/subeenchung/kaleidoscope/types"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	//Load TOML rule configuration
	m, err := filepath.Glob("./windows/*.toml")
	if err != nil {
		log.Fatalf("could not load toml rules: %v\n", err)
	}

	wrules := &types.Ruleset{}

	for _, elem := range m {
		wf := &types.WindowsRulesFormat{}
		b, err := ioutil.ReadFile(elem)
		if err != nil {
			log.Fatalf("could not read file: %v err: \n", elem, err)
		}
		err = toml.Unmarshal(b, &wf)
		if err != nil {
			log.Fatalf("err occurred while unmarshalling toml: %v\n", err)
		}
		append(&wf.Rules, wf)

	}

	for _, elem := range wf.Rules {
		fmt.Printf("%v\n", elem)
	}

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "testgrp1",
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
			if je != nil {
				log.Fatalf("could not unmarshal incoming message: %v\nmessage body: %v\n", je, string(msg.Value))
			}
			fmt.Printf("event id: %v\n", wl.Event.Code)

			//fmt.Printf("Message on %s: %s \n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer err: %v (%v)\n", err, msg)
		}
	}

	c.Close()

}
