package restApi

import (
	"net/http"
	"sample-app/exchange"
	"sample-app/orderbook"

	"github.com/gorilla/mux"
)

type Server struct {
	//ob *orderbook.OrderBook
	broker *exchange.Broker
}

func NewServer(ob *orderbook.OrderBook) *Server {
	return &Server{
		broker: exchange.NewBroker(),
	}
}

func (s *Server) Listen() {
	r := mux.NewRouter()

	// Create subrouter for /api/v1
	api := r.PathPrefix("/api/v1").Subrouter()

	// Apply middleware to the subrouter
	api.Use(validateHeadersMiddleware)

	// Define routes for the subrouter
	api.HandleFunc("/orders", s.ordersHandler)                // DELETE /GET all orders
	api.HandleFunc("/orders/limit", s.placeLimitOrderHandler) //POST /GET
	//api.HandleFunc("/orders/market", s.placeMarketOrderHandler)

	http.ListenAndServe(":8080", r)
}
