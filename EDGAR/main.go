package main

import (
	"EDGAR/getInfo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// example: https://data.sec.gov/api/xbrl/companyconcept/CIK##########/us-gaap/AccountsPayableCurrent.json

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("<Ticker>")
	}

	ticker := strings.ToUpper(os.Args[1])

	res, err := http.Get(getInfo.COMPANY_TICKERS_URL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	var tCIK *getInfo.TickerToCIK

	tCIK, err = getInfo.DownloadTickers()
	if err != nil {
		log.Fatal(err)
	}

	tickSubmission, err := getInfo.GetSubmissions(tCIK.Collection[ticker])

	jsonStr, err := json.MarshalIndent(tickSubmission, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("res.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(f, "%s\n", string(jsonStr))
}

