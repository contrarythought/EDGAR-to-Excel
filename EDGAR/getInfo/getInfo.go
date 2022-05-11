package getInfo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CompanyTicker struct {
	CIK    int    `json:"cik_str"`
	Ticker string `json:"ticker"`
	Title  string `json:"title"`
}

type CompanyCollection struct {
	Collection map[string]CompanyTicker
}

const URL = "https://www.sec.gov/files/company_tickers.json"

func (c *CompanyCollection) FromJSON(jsonStr []byte) error {
	return json.Unmarshal(jsonStr, &c.Collection)
}

func DownloadTickers() (*CompanyCollection, error) {
	var coll CompanyCollection

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := coll.FromJSON(body); err != nil {
		return nil, err
	}

	return &coll, nil
}
