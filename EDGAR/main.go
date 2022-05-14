package main

import (
	"EDGAR/MU"
	"encoding/json"
	"fmt"
	"log"
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

	var oi MU.OperatingIncomeLoss
	if err := MU.GetOperatingIncome(&oi); err != nil {
		log.Fatal(err)
	}

	s, err := json.MarshalIndent(oi, "", "	")
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%s\n", string(s))

	_, err = fmt.Fprintf(f, "%s\n", s)
	if err != nil {
		log.Fatal(err)
	}

}

/*
func readUnits(units interface{}, f *os.File) {
	switch unitType := units.(type) {
	case float64:
		//fmt.Println("float64: ", unitType)
		fmt.Fprintln(f, "float64: ", unitType)
	case string:
		fmt.Fprintln(f, "String: ", unitType)
	case []interface{}:
		fmt.Fprintln(f, "is an array: ", unitType)
		for i, u := range unitType {
			fmt.Fprint(f, i, " ")
			readUnits(u, f)
		}
	case map[string]interface{}:
		fmt.Fprintln(f, "is an object: ", unitType)
		for k, v := range unitType {
			fmt.Fprint(f, k, " ")
			readUnits(v, f)
		}
	default:
		fmt.Fprintln(f, "unknown type: ", unitType)
	}

}
*/

/*
func getBodyStr(URL string) ([]byte, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", private.USER_AGENT)
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
*/
