package main

import (
	"math/rand"
)

type Item struct {
	id    int
	name  string
	price float64
}

func getSampleItems(amount int) *[]Item {
	//sample comment
	var items []Item
	for i := 0; i < amount; i++ {
		items = append(items, Item{id: i, name: "item", price: rand.Float64()})
	}
	return &items
}
