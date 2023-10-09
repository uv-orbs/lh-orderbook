package main

import (
	"testing"
)

func Test_ModifyPriceEvery(t *testing.T) {
	//randGen := func() int { return 5 }
	stop := make(chan bool)

	market := NewMarketMock(1000)
	market.modifyPriceEvery(10)
	//p1 := market.price

	if market.price != 1005 {
		t.Errorf("Expected market.num to be 1005, but got %d", market.price)
	}

	stop <- true
}
