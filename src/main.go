package main

import (
	"log"
	"strconv"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	items := getSampleItems(10000)

	maxWorkers := 300
	queue := make(chan Item, maxWorkers)
	wg := &sync.WaitGroup{}

	for i := 0; i < maxWorkers; i++ {
		wg.Add(1)
		go worker(wg, queue)
	}

	for _, item := range *items {
		queue <- item
	}
	close(queue)
	wg.Wait()
	elapsed := time.Since(start)
	log.Printf("Rendering %s random images took %s", strconv.Itoa(len(*items)), elapsed)
	log.Print("Sleeping now for 1 hour")
	time.Sleep(3600 * time.Second)
}

func worker(wg *sync.WaitGroup, queue chan Item) {
	defer wg.Done()
	for item := range queue {
		err := renderImage(&item)
		if err != nil {
			log.Printf("Error rendering image: %v", err)
		}
	}
}
