package MU

import (
	"EDGAR/private"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const OI_URL = "https://data.sec.gov/api/xbrl/companyconcept/CIK0000723125/us-gaap/OperatingIncomeLoss.json"

type OperatingIncomeLoss struct {
	CIK         int    `json:"cik"`
	Taxonomy    string `json:"taxonomy"`
	Tag         string `json:"tag"`
	Label       string `json:"label"`
	Description string `json:"description"`
	EntityName  string `json:"entityName"`
	Units       struct {
		FileData []USD
	} `json:"units"`
}

type USD struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Val   int    `json:"val"`
	Accn  string `json:"accn"`
	FY    int    `json:"fy"`
	FP    string `json:"fp"`
	Form  string `json:"form"`
	Filed string `json:"filed"`
	Frame string `json:"frame"`
}

type OICollection struct {
	Collection map[string]OperatingIncomeLoss
}

func (oi *OICollection) GetOperatingIncome() error {
	req, err := http.NewRequest("GET", OI_URL, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", private.USER_AGENT)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resp_str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(resp_str, &oi.Collection); err != nil {
		return err
	}

	return nil
}
