package main

import (
	"EDGAR/getInfo"
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

	tickSubmission, err := getInfo.GetFacts(tCIK.Collection[ticker])

	//parseFacts(&tickSubmission.Facts.Dei)

	/*
		jsonStr, err := json.MarshalIndent(tickSubmission, "", "	")
		if err != nil {
			log.Fatal(err)
		}
	*/

	f, err := os.Create("res.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for k := range tickSubmission.Facts.Dei {
		fmt.Fprintf(f, "%s\n", k)
	}

	for k := range tickSubmission.Facts.UsGAAP {
		fmt.Fprintf(f, "%s\n", k)
	}
}

