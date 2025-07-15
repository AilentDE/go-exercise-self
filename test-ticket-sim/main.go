package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"ticket-sim/cache"
	"ticket-sim/record"
	"ticket-sim/stock"
	"time"

	"github.com/google/uuid"
)

var (
	ctx         = context.Background()
	ticketKey   = "ticket:stock"
	redisClient = cache.InitRedisClient("localhost:6379")
)

func main() {
	// init stock
	stock.PreloadTickets(redisClient, ticketKey, 150)

	http.HandleFunc("/buy", func(w http.ResponseWriter, r *http.Request) {
		// task id
		id, err := uuid.NewV7()
		if err != nil {
			log.Printf("generate uuid failed: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// buy ticket - 現在返回成功狀態和剩餘數量
		success, remainingQty := stock.HandleBuyTicket(redisClient, ticketKey, 1)
		resultRecord := record.Record{
			ID:           id.String(),
			Timestamp:    time.Now().UnixNano(),
			Result:       strconv.FormatBool(success),
			RemainingQty: remainingQty, // 新增剩餘數量
		}
		record.InsertRecord(resultRecord)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resultRecord)
	})

	http.HandleFunc("/records", func(w http.ResponseWriter, r *http.Request) {
		records := record.GetRecords()
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(records)
	})

	log.Println("server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
