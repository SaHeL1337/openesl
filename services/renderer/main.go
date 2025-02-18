package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	items "github.com/SaHeL1337/openesl/pkg/item"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type Item = items.Item

func main() {
	start := time.Now()
	nc, _ := nats.Connect("localhost:4222")
	maxWorkers := 64
	ch := make(chan jetstream.Msg, maxWorkers)
	wg := &sync.WaitGroup{}

	js, err := jetstream.New(nc)

	if err != nil {
		log.Fatalf("Error creating JetStream context: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "image3",
		Subjects: []string{"image3.*"},
	})

	if err != nil {
		log.Fatalf("Error creating JetStream stream: %v", err)
	}

	c, err := s.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "CONS",
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		log.Fatalf("Error creating JetStream consumer: %v", err)
	}
	messageCounter := 0
	cons, _ := c.Consume(func(msg jetstream.Msg) {
		ch <- msg
		messageCounter++
	})
	defer cons.Stop()

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(wg, ch)
	}

	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Rendering %s random images took %s", strconv.Itoa(messageCounter), elapsed)
	log.Print("Sleeping now for 1 hour")
	time.Sleep(3600 * time.Second)
}

func worker(wg *sync.WaitGroup, ch chan jetstream.Msg) {
	defer wg.Done()
	for msg := range ch {
		item := unmarshalMsgToItem(msg)
		log.Printf("Rendering image for item %d", item.Id)
		err := renderImage(item)
		if err != nil {
			log.Printf("Error rendering image: %v", err)
		}
	}
}

func unmarshalMsgToItem(msg jetstream.Msg) *Item {
	item := &Item{}
	err := json.Unmarshal(msg.Data(), item)
	if err != nil {
		log.Printf("Error unmarshalling item: %v", err)
	}
	msg.Ack()
	return item
}
