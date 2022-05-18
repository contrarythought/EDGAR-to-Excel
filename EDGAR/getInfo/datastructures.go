package getInfo

import "time"

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
			Street1                   string      `json:"street1"`
			Street2                   interface{} `json:"street2"`
			City                      string      `json:"city"`
			StateOrCountry            string      `json:"stateOrCountry"`
			ZipCode                   string      `json:"zipCode"`
			StateOrCountryDescription string      `json:"stateOrCountryDescription"`
		} `json:"mailing"`
		Business struct {
			Street1                   string      `json:"street1"`
			Street2                   interface{} `json:"street2"`
			City                      string      `json:"city"`
			StateOrCountry            string      `json:"stateOrCountry"`
			ZipCode                   string      `json:"zipCode"`
			StateOrCountryDescription string      `json:"stateOrCountryDescription"`
		} `json:"business"`
	} `json:"addresses"`
	Phone       string `json:"phone"`
	Flags       string `json:"flags,omitempty"`
	FormerNames []struct {
		Name string    `json:"name"`
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	}
	Filings struct {
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

type CompanyFacts struct {
	CIK        int    `json:"cik"`
	EntityName string `json:"entityName"`
	Facts      struct {
		Dei    map[string]interface{} `json:"dei"`
		UsGAAP map[string]interface{} `json:"us-gaap"`
	} `json:"facts"`
}
