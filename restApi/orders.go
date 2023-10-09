package restApi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Order struct {
	ClientOid string `json:"clientOid"`
	Side      string `json:"side"`
	Symbol    string `json:"symbol"`
	Type      string `json:"type"`
	Remark    string `json:"remark"`
	Stp       string `json:"stp"`
	TradeType string `json:"tradeType"`
	Price     string `json:"price"`
	Size      string `json:"size"`
}

func (s *Server) cancellAllOrders(symbol, tradeType string) {

	// orders := append(s.ob.GetOrderSide(orderbook.Buy).Orders(), s.ob.GetOrderSide(orderbook.Sell).Orders()...)

	// for _, element := range orders {
	// 	order, ok := element.Value.(*orderbook.Order)
	// 	if ok {
	// 		// If the type assertion was successful, print the ID
	// 		log.Println("cancelling ", order.ID())
	// 		s.ob.CancelOrder(order.ID())
	// 	} else {
	// 		// Handle the case where the element was not of type *Order
	// 		log.Println("element is not of type *Order")
	// 	}
	// }
}

// Api Handlers
func (s *Server) ordersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// handle GET request
		fmt.Fprintf(w, "GET request\n")
	case "POST":
		// handle POST request
		fmt.Fprintf(w, "POST request\n")
	case "DELETE":
		symbol := r.URL.Query().Get("symbol")
		tradeType := r.URL.Query().Get("tradeType")
		// handle DELETE request
		fmt.Fprintf(w, "Deleting order with symbol %s and tradeType %s\n", symbol, tradeType)
		//log.Print(s.ob)
		s.cancellAllOrders(symbol, tradeType)
	default:
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
func (s *Server) placeLimitOrderHandler(w http.ResponseWriter, r *http.Request) {
	var order Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Implement the logic to place a limit order here
	bts, _ := json.Marshal(order)
	log.Println(string(bts))

	sig := "0x234sig"
	s.broker.ProcMakerOrder("Maker-ID", order.Symbol, order.Side, order.Price, order.Side, sig)
}

// func (s *Server) placeMarketOrderHandler(w http.ResponseWriter, r *http.Request) {
// 	var order OrderÍ
// 	err := json.NewDecoder(r.Body).Decode(&order)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return

// 	}
// 	// Implement the logic to place a market order here
// }
