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
	println("Test ", tickSubmission.EntityName)
	f, err := os.Create("res.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	conceptList, err := getInfo.ListOfConcepts(tickSubmission)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for i := range conceptList {
			fmt.Println("[", i, "] ", conceptList[i])
		}

		var c int
		fmt.Println("Choose a concept (-1 to exit):")
		fmt.Scanf("%d\n", &c)
		if c == -1 {
			break
		}
		fmt.Println("Building a report for: ", conceptList[c])

		if err := getInfo.BuildReport(tickSubmission, conceptList[c]); err != nil {
			log.Fatal(err)
		}

		fmt.Println(conceptList[c], " report finished!")
	}

	fmt.Println("Finished!")
}
