package stock

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

func PreloadTickets(rds *redis.Client, ticketKey string, stock int) {
	err := rds.Set(ctx, ticketKey, stock, 0).Err()
	if err != nil {
		log.Printf("preload tickets failed: %v", err)
	}

	log.Printf("preload tickets success: %v", stock)
}

func HandleBuyTicket(rds *redis.Client, ticketKey string, stock int) bool {
	result, err := rds.Decr(ctx, ticketKey).Result()
	if err != nil {
		log.Printf("buy ticket failed: %v", err)
		return false
	}

	if result < 0 {
		log.Printf("ticket sold out")
		return false
	}

	log.Printf("buy ticket success: %v", result)
	return true
}
