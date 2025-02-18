package main

import (
	"log"
	"strconv"
	"time"

	items "github.com/SaHeL1337/openesl/pkg/item"
	"github.com/nats-io/nats.go"
)

type Item = items.Item

func main() {
	start := time.Now()
	nc, _ := nats.Connect("localhost:4222")

	items := items.GetSampleItems(1000)

	for _, item := range *items {
		log.Printf("Sending item %d to renderer", item.Id)
		err := nc.Publish("image.render", []byte("{\"id\": "+strconv.Itoa(item.Id)+",\"name\": \""+item.Name+"\",\"price\": "+strconv.FormatFloat(item.Price, 'f', -1, 64)+"}"))
		if err != nil {
			log.Printf("Error sending item to renderer: %v", err)
		}
		time.Sleep(0 * time.Millisecond)
	}

	elapsed := time.Since(start)
	log.Printf("Sending %s random items took %s", strconv.Itoa(len(*items)), elapsed)
	log.Print("Sleeping now for 1 hour")
	time.Sleep(3600 * time.Second)
}
