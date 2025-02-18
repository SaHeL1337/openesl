package main

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	start := time.Now()
	nc, _ := nats.Connect("localhost:4222")
	maxWorkers := 64
	ch := make(chan *nats.Msg, maxWorkers)
	wg := &sync.WaitGroup{}
	sub, _ := nc.ChanSubscribe("image.render", ch)
	items := []Item{}

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(wg, ch)
	}

	wg.Wait()
	sub.Unsubscribe()
	elapsed := time.Since(start)
	log.Printf("Rendering %s random images took %s", strconv.Itoa(len(items)), elapsed)
	log.Print("Sleeping now for 1 hour")
	time.Sleep(3600 * time.Second)
}

func worker(wg *sync.WaitGroup, ch chan *nats.Msg) {
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

func unmarshalMsgToItem(msg *nats.Msg) *Item {
	item := &Item{}
	err := json.Unmarshal(msg.Data, item)
	if err != nil {
		log.Printf("Error unmarshalling item: %v", err)
	}
	return item
}
