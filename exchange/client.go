package exchange

import "sample-app/orderbook"

// Define the Client struct
type client struct {
	// Add fields as needed
	id            string
	OpenOrders    map[string]Order // open for
	PendingOrders map[string]Order // pending onchain qompletion
}

// Define the newBroker function
func newClient(id string) *client {
	return &client{
		id:            id,
		OpenOrders:    make(map[string]Order),
		PendingOrders: make(map[string]Order),
	}
}

func (c *client) checkFunds() bool {
	return true
}

func (c *client) saveOrder(order *orderbook.Order) bool {
	return true
}
