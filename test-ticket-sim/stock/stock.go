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

func HandleBuyTicket(rds *redis.Client, ticketKey string, stock int) (bool, int64) {
	// 使用 Lua 腳本保證原子性
	script := `
		local key = KEYS[1]
		local current_stock = redis.call('GET', key)
		
		if tonumber(current_stock) <= 0 then
			return {0, 0}  -- 失敗，剩餘數量為 0
		end
		
		local new_stock = redis.call('DECR', key)
		return {1, new_stock}  -- 成功，返回新的剩餘數量
	`

	result, err := rds.Eval(ctx, script, []string{ticketKey}).Result()
	if err != nil {
		log.Printf("buy ticket failed: %v", err)
		return false, -1
	}

	results := result.([]interface{})
	success := results[0].(int64) == 1
	remainingQty := results[1].(int64)

	if success {
		log.Printf("buy ticket success: %v", remainingQty)
	} else {
		log.Printf("ticket sold out")
	}

	return success, remainingQty
}
