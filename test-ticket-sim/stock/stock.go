package stock

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
)

func PreloadTickets(rds *redis.Client, ticketKey string, stock int) error {
	// 檢查是否已經有庫存
	_, err := rds.Get(ctx, ticketKey).Result()
	if err == nil {
		// 已經有庫存，直接返回
		return nil
	}

	// 如果 key 不存在或庫存為 0，則初始化庫存
	for i := 0; i < 10; i++ {
		ok, err := rds.SetNX(ctx, ticketKey, stock, time.Second*10).Result()
		if err != nil {
			log.Printf("setNX failed: %v", err)
			continue
		}

		if ok {
			log.Printf("preload tickets success: %v", stock)
			return nil
		}

		// 如果 SetNX 失敗，檢查是否其他進程已經設置了庫存
		time.Sleep(time.Millisecond * 10)
		_, err = rds.Get(ctx, ticketKey).Result()
		if err == nil {
			return nil
		}

	}

	return errors.New("preload tickets failed")
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
