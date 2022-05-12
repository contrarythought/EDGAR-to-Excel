package MU

// "https://data.sec.gov/api/xbrl/companyconcept/CIK0000723125/us-gaap/OperatingIncomeLoss.json"
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

func (oi *OperatingIncomeLoss) GetOperatingIncome() error {
	
}
