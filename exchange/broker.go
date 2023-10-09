package exchange

import (
	"fmt"
	"log"
	"sample-app/orderbook"

	"github.com/shopspring/decimal"
)

// Define the Broker type
type Broker struct {
	clients map[string]*client
	pairs   map[string]*Pair
	index   uint
}

// Define the newBroker function
func NewBroker() *Broker {
	return &Broker{
		clients: make(map[string]*client),
		pairs:   make(map[string]*Pair),
		index:   0,
	}
}

func (b *Broker) clientId2BrokerId(pairId, cid string) string {
	b.index++
	return fmt.Sprintf("%d-%s-%s", b.index, pairId, cid)
}

// Define the createOrder method for the Broker type (MarketMaker side)
func (b *Broker) ProcMakerOrder(cid, pairId, side, price, size, signature string) error {
	// Add your order creation logic here

	client, ok := b.clients[cid]
	if !ok {
		return fmt.Errorf("clientID is not found: %s", cid)
	}

	if !client.checkFunds() {
		return fmt.Errorf(fmt.Sprintf("funds are not sufficient: %s", cid))
	}

	pair, ok := b.pairs[pairId]
	if !ok {
		return fmt.Errorf(fmt.Sprintf("pair id not found: ", pairId))
	}
	dcPrice, err := decimal.NewFromString(price)
	if err != nil {
		return err
	}
	dcSize, err := decimal.NewFromString(size)
	if err != nil {
		return err
	}
	// side
	obSide := orderbook.Buy
	if side == "sell" {
		obSide = orderbook.Sell
	}
	brokerId := b.clientId2BrokerId(cid, pairId)
	done, partial, partialQuantityProcessed, err := pair.ob.ProcessLimitOrder(obSide, brokerId, dcPrice, dcSize)
	log.Println(done)
	log.Println(partial)
	log.Println(partialQuantityProcessed)
	if err != nil {
		return err
	}
	return nil
}
