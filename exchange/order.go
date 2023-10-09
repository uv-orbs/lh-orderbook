package exchange

// Define the Pair struct
type Order struct {
	// Add fields as needed
	id        string
	price     string
	size      string
	signature string // EIP 712
	pending   bool
}

// Define the newBroker function
func newOrder(id, price, size string, pending bool) *Order {
	return &Order{
		id:      id,
		price:   price,
		size:    size,
		pending: false,
	}
}
