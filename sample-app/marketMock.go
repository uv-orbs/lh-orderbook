package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sample-app/restApi"
	"time"
)

type MarketMock struct {
	price      uint
	orderIndex uint
}

func NewMarketMock(price uint) *MarketMock {
	return &MarketMock{
		price:      1000,
		orderIndex: 0,
	}
}

func httpMethod(method, url string) {
	// Create a new request using http
	req, err := http.NewRequest("DELETE", url, nil)

	if err != nil {
		fmt.Println("Error creating request: ", err)
		return
	}
	setHeaders(req)
	//req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	// Send the request using the http Client
	//client := &http.Client{}
	//resp, err := client.Do(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error sending request: ", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body: ", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}

func (m *MarketMock) onNewPrice() {
	// cancel old olders
	httpMethod("DELETE", "http://localhost:8080/api/v1/orders")

	// send 3 ask above the price
	for i := 0; i < 3; i++ {
		u := uint(i) + 1
		m.placeOrder("sell", m.price+u, u)
	}

	// send 3 bids below the price
	for i := 0; i < 3; i++ {
		u := uint(i) + 1
		m.placeOrder("buy", m.price-u, u)
	}

}

const apiKey = "1afa7e8f1d307e934adb042b2e5568639c4846d09440f7551cefbfb0a2121db0" // key for "your-pass-phrase"
func setHeaders(req *http.Request) {
	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("API-PASSPHRASE", "your-pass-phrase")
	req.Header.Set("API-KEY", apiKey)
}

func (m *MarketMock) placeOrder(side string, price uint, size uint) {
	m.orderIndex += 1
	oid := fmt.Sprintf("maker-%d-%s-%d-%d", m.orderIndex, side, price, size)
	order := restApi.Order{
		ClientOid: oid,
		Side:      side,
		Symbol:    "ETH-BTC",
		Type:      "limit",
		Remark:    fmt.Sprintf("remark%d", rand.Int()),
		Stp:       "CN",
		TradeType: "TRADE",
		Price:     fmt.Sprintf("%.2f", float32(price)),
		Size:      fmt.Sprintf("%.2f", float32(size)),
	}

	orderJson, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error marshalling order:", err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/api/v1/orders/limit", bytes.NewBuffer(orderJson))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	setHeaders(req)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending order:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Order sent, response status:", resp.Status)
}
func (m *MarketMock) modifyPriceEvery(sec uint) {
	// Initialize the uint number with 1000
	m.price = 1000

	// Setup a ticker that fires every minute
	ticker := time.NewTicker(time.Second * time.Duration(sec))
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Generate a random number from -10 to +10
			randNum := rand.Intn(21) - 10

			// If the random number is negative and its absolute value is greater than num,
			// then reset num to 0, else add the random number to num
			if randNum < 0 && uint(-randNum) > m.price {
				m.price = 1000
			} else {
				m.price = m.price + uint(randNum)
				m.onNewPrice()
			}

			fmt.Printf("Current price is: %d\n", m.price)
		}
	}
}
