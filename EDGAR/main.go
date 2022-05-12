package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const CIK_SIZE = 10

type TickerInfo struct {
	tickerCIK map[string]string
}

// example: https://data.sec.gov/api/xbrl/companyconcept/CIK##########/us-gaap/AccountsPayableCurrent.json

func main() {
	/*
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
			cikStr := getInfo.CIKToString(v.CIK)
			t.tickerCIK[v.Ticker] = cikStr
		}

		fmt.Printf("%s\n", t.tickerCIK["MU"])
	*/

	f, err := os.Create("res.txt")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := getBodyStr("https://data.sec.gov/api/xbrl/companyconcept/CIK0000723125/us-gaap/OperatingIncomeLoss.json")
	if err != nil {
		log.Fatal(err)
	}

	/*
		resp, err = json.MarshalIndent(resp, "", "	")
		if err != nil {
			log.Fatal(err)
		}
	*/
	fmt.Fprintf(f, "%s\n", resp)

}

func getBodyStr(URL string) ([]byte, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "athorpe624@gmail.com")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
