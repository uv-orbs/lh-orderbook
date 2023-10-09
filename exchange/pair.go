package exchange

import (
	"fmt"
	"sample-app/orderbook"
)

// Define the Pair struct
type Pair struct {
	// Add fields as needed
	id      string
	symbolA string
	symbolB string
	ob      *orderbook.OrderBook
}

// Define the newBroker function
func newPair(symbolA, symbolB string) *Pair {
	id := fmt.Sprintf("%s-%s", symbolA, symbolB)
	ob := orderbook.NewOrderBook()
	return &Pair{
		id:      id,
		symbolA: symbolA,
		symbolB: symbolA,
		ob:      ob,
	}
}
