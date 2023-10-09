package main

import (
	"fmt"
	"log"
	"sample-app/orderbook"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func Test_Main(t *testing.T) {
	ob := orderbook.NewOrderBook()

	// buy orders
	quantity := decimal.New(2, 0)
	for i := 5; i < 10; i = i + 1 {
		done, partial, partialQty, err := ob.ProcessLimitOrder(orderbook.Buy, fmt.Sprintf("buy-%d", i), quantity, decimal.New(int64(i), 0))
		if len(done) != 0 {
			log.Fatal("OrderBook failed to process limit order (done is not empty)")
		}
		if partial != nil {
			log.Fatal("OrderBook failed to process limit order (partial is not empty)")
		}
		if partialQty.Sign() != 0 {
			log.Fatal("OrderBook failed to process limit order (partialQty is not zero)")
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	// sell orders
	for i := 10; i < 15; i = i + 1 {
		done, partial, partialQty, err := ob.ProcessLimitOrder(orderbook.Sell, fmt.Sprintf("sell-%d", i), quantity, decimal.New(int64(i), 0))
		if len(done) != 0 {
			log.Fatal("OrderBook failed to process limit order (done is not empty)")
		}
		if partial != nil {
			log.Fatal("OrderBook failed to process limit order (partial is not empty)")
		}
		if partialQty.Sign() != 0 {
			log.Fatal("OrderBook failed to process limit order (partialQty is not zero)")
		}
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Print(ob)

	// taker in 100 + 110 and print again
	price := decimal.New(int64(100), 0)
	log.Println("buy 100")
	quantity = decimal.New(3, 0)
	done, partial, partialQuantityProcessed, err := ob.ProcessLimitOrder(orderbook.Buy, fmt.Sprintf("take-%d", price), quantity, price)
	log.Println(done, partial, partialQuantityProcessed, err)

	// log.Println("buy 110")
	// price = decimal.New(int64(110), 0)
	// done, partial, partialQuantityProcessed, err = ob.ProcessLimitOrder(orderbook.Buy, fmt.Sprintf("take-%d", price), quantity, price)
	// log.Println(done, partial, partialQuantityProcessed, err)

	log.Print("+++++++++++++++++++++++++++++++++++++++++++++++++")
	log.Print(ob)

	for i := 0; i < 200; i = i + 1 {
		log.Print("+++++++++++++++++++++++++++++++++++++++++++++++++")
		randomTrade(ob)
		log.Print(ob)
		time.Sleep(1 * time.Second)
	}

	// data := []*OB.Order{
	// 	OB.NewOrder("one", OB.Buy, decimal.New(11, -1), decimal.New(11, 1), time.Now().UTC()),
	// 	OB.NewOrder("two", OB.Buy, decimal.New(22, -1), decimal.New(22, 1), time.Now().UTC()),
	// 	OB.NewOrder("three", OB.Sell, decimal.New(33, -1), decimal.New(33, 1), time.Now().UTC()),
	// 	OB.NewOrder("four", OB.Sell, decimal.New(44, -1), decimal.New(44, 1), time.Now().UTC()),
	// }

	// result, _ := json.Marshal(data)
	// fmt.Println(string(result))

	// data = []*OB.Order{}

	// _ = json.Unmarshal(result, &data)
	// fmt.Println(data)

	// err := json.Unmarshal([]byte(`[{"side":"fake"}]`), &data)
	// if err == nil {
	// 	log.Fatalf("can unmarshal unsupported value")
	// }
}
