package main

import (
	"EDGAR/getInfo"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const CIK_SIZE = 10

type TickerInfo struct {
	tickerCIK map[string]string
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

	var t TickerInfo = TickerInfo{make(map[string]string)}
	for _, v := range c.Collection {
		cikStr := CIKToString(v.CIK)
		t.tickerCIK[v.Ticker] = cikStr
	}

	fmt.Printf("%s\n", t.tickerCIK["BNED"])
	fmt.Printf("%s\n", t.tickerCIK["AAPL"])
	fmt.Printf("%s\n", t.tickerCIK["ALOT"])
}

// try with built-in functions?
func CIKToString(CIK int) string {
	var ret []byte
	cikStr := strconv.Itoa(CIK)
	cikLen := len(cikStr)

	if cikLen < CIK_SIZE {
		ret = make([]byte, CIK_SIZE)
		offset := CIK_SIZE - cikLen
		var i int
		// prepend return CIK with '0' if cikLen < 10 
		for i = 0; i < offset; i++ {
			ret[i] = '0'
		}

		// copy over the rest of the CIK string
		for j := 0 ; i < CIK_SIZE; i, j = i + 1, j + 1 {
			ret[i] = cikStr[j]
		}

		return string(ret)
	}	

	return cikStr
}
