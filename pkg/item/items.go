package item

import (
	"math/rand"
)

type Item struct {
	Id    int
	Name  string
	Price float64
}

func GetSampleItems(amount int) *[]Item {
	//sample comment
	var items []Item
	for i := 0; i < amount; i++ {
		items = append(items, Item{Id: i, Name: "item", Price: rand.Float64()})
	}
	return &items
}
