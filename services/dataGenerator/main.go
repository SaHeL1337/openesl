package main

import (
	"context"
	"log"
	"strconv"
	"time"

	items "github.com/SaHeL1337/openesl/pkg/item"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type Item = items.Item

func main() {
	start := time.Now()
	nc, _ := nats.Connect("localhost:4222")

	items := items.GetSampleItems(1000)

	js, _ := jetstream.New(nc)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	for _, item := range *items {
		js.Publish(ctx, "image3.render", []byte("{\"id\": "+strconv.Itoa(item.Id)+",\"name\": \""+item.Name+"\",\"price\": "+strconv.FormatFloat(item.Price, 'f', -1, 64)+"}"))
		log.Printf("Sending item %d to renderer", item.Id)
	}

	elapsed := time.Since(start)
	log.Printf("Sending %s random items took %s", strconv.Itoa(len(*items)), elapsed)
	log.Print("Sleeping now for 1 hour")
	time.Sleep(3600 * time.Second)
}
