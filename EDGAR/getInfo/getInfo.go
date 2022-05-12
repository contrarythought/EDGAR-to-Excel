package getInfo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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
const CIK_SIZE = 10

func (c *CompanyCollection) FromJSON(jsonStr []byte) error {
	return json.Unmarshal(jsonStr, &c.Collection)
}

func CIKToString(CIK int) string {
	var ret []byte
	cikStr := strconv.Itoa(CIK)
	cikLen := len(cikStr)

	if cikLen < CIK_SIZE {
		ret = make([]byte, CIK_SIZE)
		offset := CIK_SIZE - cikLen
		var i int
		// prepend return CIK with '0' if cikLen < 10
		for i = 0; i < offset; i++ {
			ret[i] = '0'
		}

		// copy over the rest of the CIK string
		for j := 0; i < CIK_SIZE; i, j = i+1, j+1 {
			ret[i] = cikStr[j]
		}

		return string(ret)
	}

	return cikStr
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
