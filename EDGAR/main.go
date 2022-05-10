package main

import (
	"EDGAR/getInfo"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get(getInfo.URL)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	f, err := os.Create("res.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	bodyString, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Failed to read into bodyString: %s", err)
	}

	var c getInfo.CompanyCollection
	if err := c.FromJSON(bodyString); err != nil {
		log.Fatal(err)
	}

	for _, v := range c.Collection {
		fmt.Fprintf(f, "%s %v\n", v.Ticker, v.CIK)
	}
}
