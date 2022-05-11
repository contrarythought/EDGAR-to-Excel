package main

import (
	"EDGAR/getInfo"
	"fmt"
	"log"
	"net/http"
)

type TickerInfo struct {
	tickerCIK map[string]getInfo.CompanyTicker
}

func main() {
	res, err := http.Get(getInfo.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var c *getInfo.CompanyCollection

	c, err = getInfo.DownloadTickers()
	if err != nil {
		log.Fatal(err)
	}

	var t TickerInfo = TickerInfo{make(map[string]getInfo.CompanyTicker)}
	for _, v := range c.Collection {
		t.tickerCIK[v.Ticker] = v
	}

	fmt.Printf("%s\n", t.tickerCIK["BNED"].Title)
	fmt.Printf("%d\n", t.tickerCIK["AAPL"].CIK)
	fmt.Printf("%s\n", t.tickerCIK["ALOT"].Title)
}
