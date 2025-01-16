package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	items := getSampleItems(1000)

	for i, item := range *items {
		err := renderImage(&item)
		if err != nil {
			log.Printf("Error rendering image: %v", err)
		}
		if i%100 == 0 {
			elapsed := time.Since(start)
			fmt.Printf("Rendering %s images in %s \n", strconv.Itoa(i), elapsed)
		}
	}
	elapsed := time.Since(start)
	log.Printf("Rendering %s random images took %s", strconv.Itoa(len(*items)), elapsed)
}
