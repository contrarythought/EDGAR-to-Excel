package getInfo

import (
	"encoding/json"
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

func DownloadTickers() (map[string]CompanyTicker, error) {
	ret := make(map[string]CompanyTicker)

	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var t CompanyTicker
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		return nil, err
	}

	return ret, nil
}
