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

	f, err := os.Create("res.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	conceptList, err := getInfo.ListOfConcepts(tickSubmission)
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range conceptList {
		fmt.Println("[", i, "]", s)
	}

	con, err := getInfo.GetConcept(tickSubmission, "Revenues")
	if err != nil {
		log.Fatal(err)
	}

	for i := range con.Units.USD {
		fmt.Fprintf(f, "Start: %s End: %s --- %d\n", con.Units.USD[i].Start, con.Units.USD[i].End, con.Units.USD[i].Val)
	}

}
