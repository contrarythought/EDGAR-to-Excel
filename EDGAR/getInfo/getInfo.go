package getInfo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type CompanyTicker struct {
	CIK    int    `json:"cik_str"`
	Ticker string `json:"ticker"`
	Title  string `json:"title"`
}

type CompanyCollection struct {
	Collection map[string]CompanyTicker
}

type TickerToCIK struct {
	Collection map[string]string
}

type Submissions struct {
	CIK                               string   `json:"cik"`
	EntityType                        string   `json:"entityType"`
	SIC                               string   `json:"sic"`
	SICDescription                    string   `json:"sicDescription"`
	InsiderTransactionForOwnerExists  int      `json:"insiderTransactionForOwnerExists"`
	InsiderTransactionForIssuerExists int      `json:"insiderTransactionForIssuerExists"`
	Name                              string   `json:"name"`
	Tickers                           []string `json:"tickers"`
	Exchanges                         []string `json:"exchanges"`
	EIN                               string   `json:"ein"`
	Description                       string   `json:"description,omitempty"`
	Website                           string   `json:"website,omitempty"`
	InvestorWebsite                   string   `json:"investorWebsite,omitempty"`
	Category                          string   `json:"category"`
	FiscalYearEnd                     string   `json:"fiscalYearEnd"`
	StateOfIncorporation              string   `json:"stateOfIncorporation"`
	StateOfIncorporationDescription   string   `json:"stateOfIncorporationDescription"`
	Addresses                         struct {
		Mailing struct {
			Street1                   string `json:"street1"`
			Street2                   string `json:"street2"`
			City                      string `json:"city"`
			StateOrCountry            string `json:"stateOrCountry"`
			ZipCode                   string `json:"zipCode"`
			StateOrCountryDescription string `json:"stateOrCountryDescription"`
		} `json:"mailing"`
		Business struct {
			Street1                   string `json:"street1"`
			Street2                   string `json:"street2"`
			City                      string `json:"city"`
			StateOrCountry            string `json:"stateOrCountry"`
			ZipCode                   string `json:"zipCode"`
			StateOrCountryDescription string `json:"stateOrCountryDescription"`
		} `json:"business"`
	} `json:"addresses"`
	Phone       string   `json:"phone"`
	Flags       string   `json:"flags,omitempty"`
	FormerNames []string `json:"formerNames,omitifempty"`
	Filings     struct {
		Recent struct {
			AccessionNumber    []string    `json:"accessionNumber"`
			FilingDate         []string    `json:"filingDate"`
			ReportDate         []string    `json:"reportDate"`
			AcceptanceDateTime []time.Time `json:"acceptanceDateTime"`
			Act                []string    `json:"act"`
			Form               []string    `json:"form"`
			FileNumber         []string    `json:"fileNumber"`
			FilmNumber         []string    `json:"filmNumber"`
			Items              []string    `json:"items"`
			Size               []int       `json:"size"`
			IsXBRL             []int       `json:"isXBRL"`
			IsInlineXBRL       []int       `json:"isInlineXBRL"`
			PrimaryDocument    []string    `json:"primaryDocument"`
		} `json:"recent"`
		Files []struct {
			Name        string `json:"name"`
			FilingCount int    `json:"filingCount"`
			FilingFrom  string `json:"filingFrom"`
			FilingTo    string `json:"filingTo"`
		} `json:"files"`
	} `json:"filings"`
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

func initTickerToCIK() *TickerToCIK {
	return &TickerToCIK{make(map[string]string)}
}

func DownloadTickers() (*CompanyCollection, *TickerToCIK, error) {
	coll := &CompanyCollection{}

	resp, err := http.Get(URL)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err := coll.FromJSON(body); err != nil {
		return nil, nil, err
	}

	ticker_to_CIK := initTickerToCIK()
	for k, v := range coll.Collection {
		ticker_to_CIK.Collection[k] = CIKToString(v.CIK)
	}

	return coll, ticker_to_CIK, nil
}
