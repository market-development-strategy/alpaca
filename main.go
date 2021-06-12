package main

import (
	"fmt"
	"os"

	"github.com/alpacahq/alpaca-trade-api-go/alpaca"
	"github.com/alpacahq/alpaca-trade-api-go/common"
	"github.com/alpacahq/alpaca-trade-api-go/stream"
)

func main() {

	os.Setenv(common.EnvApiKeyID, os.Getenv("alpaca_key_id"))
	os.Setenv(common.EnvApiSecretKey, os.Getenv("alpaca_secret_key"))


	if err := stream.Register(alpaca.TradeUpdates, tradeHandler); err != nil {
		panic(err)
	}

	if err := stream.Register("Q.AAPL", quoteHandler); err != nil {
		panic(err)
	}

	select {}
}

func tradeHandler(msg interface{}) {
	tradeupdate := msg.(alpaca.TradeUpdate)
	fmt.Printf("%s event received for order %s.\n", tradeupdate.Event, tradeupdate.Order.ID)
}

func quoteHandler(msg interface{}) {
	quote := msg.(alpaca.StreamQuote)

	fmt.Println(quote.Symbol, quote.BidPrice, quote.BidSize, quote.AskPrice, quote.AskSize)
}
