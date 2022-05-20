package main

import (
	"EDGAR/getInfo"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
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

	var w sync.WaitGroup
	var cnt int
	for {
		if cnt > 5 {
			fmt.Println("Pausing for 40 seconds to ease load on server")
			time.Sleep(40 * time.Second)
			cnt = 0
			fmt.Println("Back again!")
		}

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

		w.Add(1)
		cnt++
		fmt.Println("Reports being built: ", cnt)
		go func(index *int) {
			if err := BuildReport(tickSubmission, conceptList[*index]); err != nil {
				log.Fatal(err)
			}

			w.Done()
			cnt--
			fmt.Println(conceptList[*index], " report finished")
		}(&c)
	}
	w.Wait()

	fmt.Println("Finished!")
}

func BuildReport(facts *getInfo.CompanyFacts, concept string) error {
	con, err := getInfo.GetConcept(facts, concept)
	if err != nil {
		return err
	}

	f, err := os.Create(concept + "Report.txt")
	if err != nil {
		return err
	}

	if len(con.Units.USD) > 0 {
		for _, v := range con.Units.USD {
			fmt.Fprintln(f, v.Form, " Start: ", v.Start, " End: ", v.End,
				" ----> ", v.Val)
		}
	} else if len(con.Units.EUR) > 0 {
		for _, v := range con.Units.EUR {
			fmt.Fprintln(f, v.Form, " Start: ", v.Start, " End: ", v.End,
				" ----> ", v.Val)
		}
	} else {
		return fmt.Errorf("unsupported accounting standard")
	}

	return nil
}
