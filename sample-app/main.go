package main

import (
	"fmt"
	"math/rand"
	"sample-app/orderbook"
	"sample-app/restApi"
	"time"

	"github.com/shopspring/decimal"
)

func gnrtRndOrders(plvl []*orderbook.PriceLevel, side orderbook.Side) []*orderbook.Order {
	minPrice := int(plvl[len(plvl)-1].Price.IntPart())
	orders := []*orderbook.Order{}
	for index, element := range plvl {
		id := fmt.Sprintf("rnd-%s-%d", side, index)
		price := int(element.Price.IntPart())
		price = rand.Intn(price-minPrice+1) + minPrice + 1
		qnt := int(element.Quantity.IntPart()/2) + 1
		qnt = rand.Intn(qnt) + 1
		//log.Println(index, price, qnt)
		order := orderbook.NewOrder(id, side, decimal.New(int64(qnt), 0), decimal.New(int64(price), 0), time.Now())
		orders = append(orders, order)
	}

	// Fill MaxBid and MinAsk
	if false {
		var lastLvl *orderbook.PriceLevel
		var opSide orderbook.Side
		if side == orderbook.Sell {
			lastLvl = plvl[len(plvl)-1]
			opSide = orderbook.Buy
		} else {
			lastLvl = plvl[0]
			opSide = orderbook.Sell
		}
		id := fmt.Sprintf("fill-%s-%d", opSide, len(plvl))
		qnt := int(lastLvl.Quantity.IntPart())
		qnt = rand.Intn(qnt) + qnt
		order := orderbook.NewOrder(id, opSide, decimal.New(int64(qnt), 0), lastLvl.Price, time.Now())
		orders = append(orders, order)
	}

	return orders
}

func randomTrade(ob *orderbook.OrderBook) {
	asks, bids := ob.Depth()

	// append to OB 10%
	// rndAsks := gnrtRndOrders(asks, orderbook.Sell)
	// rndBids := gnrtRndOrders(bids, orderbook.Buy)

	// asks
	max := int(asks[0].Price.IntPart())
	min := int(asks[len(asks)-1].Price.IntPart())
	rndPrice := rand.Intn(max-min) + min + 1
	rndID := rand.Intn(10000) + 10000
	id := fmt.Sprintf("rnd-%s-%d-%d", orderbook.Sell, rndID, rndPrice)
	ob.ProcessLimitOrder(orderbook.Sell, id, decimal.New(int64(1), 0), decimal.New(int64(rndPrice), 0))

	// bids
	max = int(bids[0].Price.IntPart()) + 1
	min = int(bids[len(bids)-1].Price.IntPart())
	rndPrice = rand.Intn(max-min) + min
	rndID = rand.Intn(10000) + 10000
	id = fmt.Sprintf("rnd-%s-%d-%d", orderbook.Buy, rndID, rndPrice)
	ob.ProcessLimitOrder(orderbook.Buy, id, decimal.New(int64(1), 0), decimal.New(int64(rndPrice), 0))

	// play rand orders
	// for _, e := range rndAsks {
	// 	ob.ProcessLimitOrder(e.Side(), e.ID(), e.Quantity(), e.Price())
	// }
	// for _, e := range rndBids {
	// 	ob.ProcessLimitOrder(e.Side(), e.ID(), e.Quantity(), e.Price())
	// }

	//log.Println(asks[0].Price.String())
	//log.Println(.String())
	//log.Println(bids[0].Price.String())
	//log.Println(bids[len(bids)-1].Price.String())

}
func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// name := "Go Developers"
	// fmt.Println("Azure for", name)
	// ob := orderbook.NewOrderBook()
	// log.Print(ob.String())

	initPrice := uint(1000)
	marketMock := NewMarketMock(initPrice)
	// start market making
	seconds := 3
	go marketMock.modifyPriceEvery(uint(seconds))
	// start http server
	ob := orderbook.NewOrderBook()
	srv := restApi.NewServer(ob)
	srv.Listen()

}
