package stored_data

import (
	"fmt"
	"time"
)

type Item struct {
	Id int
	name string
	price float64
	createdAt time.Time
}

func CreateItems(count int) *[]Item {
	items := make([]Item, count)
	for i := 0; i < int(count); i++ {
		items[i] = Item{i, fmt.Sprintf("item %d", i), 1.0 * float64(i), time.Now()}
	}

	return &items
}