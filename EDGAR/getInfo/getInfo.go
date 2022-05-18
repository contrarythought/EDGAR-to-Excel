package getInfo

import (
	"EDGAR/private"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	COMPANY_TICKERS_URL   = "https://www.sec.gov/files/company_tickers.json"
	SUBMISSION_BASE_URL   = "https://data.sec.gov/submissions/CIK"
	COMPANYFACTS_BASE_URL = "https://data.sec.gov/api/xbrl/companyfacts/CIK"
	CIK_SIZE              = 10
)

func (c *CompanyCollection) FromJSON(jsonStr []byte) error {
	return json.Unmarshal(jsonStr, &c.Collection)
}

func customGET(URL string, userAgent string) ([]byte, error) {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	ret, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

/*
	input: CIK number string that is used to set up URL for the GET request for the submissions data
	return: pointer to CompanyFacts struct for the inputted CIK, as well as an error
*/
func GetFacts(CIK string) (*CompanyFacts, error) {
	var ret CompanyFacts
	url_to_send := COMPANYFACTS_BASE_URL + CIK + ".json"
	body, err := customGET(url_to_send, private.USER_AGENT)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

/*
	input: CIK number string that is used to set up URL for the GET request for the submissions data
	return: pointer to Submissions struct for the inputted CIK, as well as an error
*/
func GetSubmissions(CIK string) (*Submissions, error) {
	var ret = Submissions{}
	url_to_send := SUBMISSION_BASE_URL + CIK + ".json"
	body, err := customGET(url_to_send, private.USER_AGENT)
	if err != nil {
		return nil, err
	}

	// unmarshal into Submissions struct
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
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

func initTickerToCIK() *TickerToCIK {
	return &TickerToCIK{make(map[string]string)}
}

func DownloadTickers() (*TickerToCIK, error) {
	coll := &CompanyCollection{}

	resp, err := http.Get(COMPANY_TICKERS_URL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := coll.FromJSON(body); err != nil {
		return nil, err
	}

	ticker_to_CIK := initTickerToCIK()
	for _, v := range coll.Collection {
		ticker_to_CIK.Collection[v.Ticker] = CIKToString(v.CIK)
	}

	return ticker_to_CIK, nil
}
