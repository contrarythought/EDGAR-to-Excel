package getInfo

import (
	"EDGAR/private"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

const (
	COMPANY_TICKERS_URL     = "https://www.sec.gov/files/company_tickers.json"
	SUBMISSION_BASE_URL     = "https://data.sec.gov/submissions/CIK"
	COMPANYFACTS_BASE_URL   = "https://data.sec.gov/api/xbrl/companyfacts/CIK"
	COMPANYCONCEPT_BASE_URL = "https://data.sec.gov/api/xbrl/companyconcept/CIK"
	CIK_SIZE                = 10
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
	fmt.Printf("Sending %s\n", url_to_send)
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

// returns the accounting standard associated with a company
func GetAccountingStd(facts *CompanyFacts) (string, error) {
	if facts == nil {
		return "", fmt.Errorf("empty CompanyFacts pointer")
	}

	if len(facts.Facts.UsGAAP) > 0 {
		return "us-gaap", nil
	} else if len(facts.Facts.IFRSFull) > 0 {
		return "ifrs-full", nil
	} else {
		return "", fmt.Errorf("cannot validate accounting standards - check JSON file")
	}
}

// returns a string array of the concepts associated with a company
func ListOfConcepts(facts *CompanyFacts) ([]string, error) {
	if facts == nil {
		return nil, fmt.Errorf("empty CompanyFacts pointer")
	}

	accounting, err := GetAccountingStd(facts)
	if err != nil {
		return nil, err
	}

	var ret []string

	if strings.Compare(accounting, "us-gaap") == 0 {
		for k := range facts.Facts.UsGAAP {
			str := fmt.Sprintf("%v", facts.Facts.UsGAAP[k])
			if strings.Contains(str, "Deprecated") {
				ret = append(ret, k+"_DEPRECATED"+" -- "+str)
			} else {
				ret = append(ret, k)
			}
		}
	} else if strings.Compare(accounting, "ifrs-full") == 0 {
		for k := range facts.Facts.IFRSFull {
			str := fmt.Sprintf("%v", facts.Facts.IFRSFull[k])
			if strings.Contains(str, "Deprecated") {
				ret = append(ret, k+"_DEPRECATED"+" -- "+str)
			} else {
				ret = append(ret, k)
			}
		}
	}

	sort.Strings(ret)
	return ret, nil
}

/*
	input: CompanyFacts pointer to obtain the accounting standard, and the concept string that you want to view
	output: CompanyConcept pointer containing concept info, and an error
*/
func GetConcept(facts *CompanyFacts, concept string) (*CompanyConcept, error) {
	accounting, err := GetAccountingStd(facts)
	if err != nil {
		return nil, err
	}

	url_to_send := COMPANYCONCEPT_BASE_URL + CIKToString(facts.CIK) + "/" + accounting + "/" + concept + ".json"
	fmt.Printf("sending: %s\n", url_to_send)
	body, err := customGET(url_to_send, private.USER_AGENT)
	if err != nil {
		return nil, err
	}

	var ret = CompanyConcept{}
	if err = json.Unmarshal(body, &ret); err != nil {
		return nil, err
	}

	return &ret, nil
}

// increments row in an excel worksheet
func incrRow(index int) string {
	ret := "A" + strconv.Itoa(index+1)
	return ret
}

// returns col string - thank you https://stackoverflow.com/questions/3302857/algorithm-to-get-the-excel-like-column-name-of-a-number
func getCol(alpha_idx int, row_idx int) string {
	ret := ""
	for ; alpha_idx >= 0; alpha_idx = (alpha_idx / 26) - 1 {
		ret = string(rune(('A' + alpha_idx%26))) + ret
	}
	return ret + strconv.Itoa(row_idx+1)
}

// TODO
func sortByFilingDates(con *CompanyConcept) {
	if len(con.Units.USD) > 0 {
		sort.Sort(USDArray(con.Units.USD))
	}
}

// TODO
func printUnit(unit interface{}, row *int) {
	switch unit.(type) {
	case USD:

	}
}

// TODO
func printConcept(con *CompanyConcept, f *excelize.File, filename string, rowIter int) (int, error) {
	axis := ""
	var err error
	var used = make(map[string]bool)

	if len(con.Units.USD) > 0 {
		// print concept type
		axis = getCol(0, rowIter)
		f.SetCellValue("Sheet1", axis, con.Tag)

		rowIter++
		len_USD := len(con.Units.USD)

		/** TODO **/
		// try to wrap print into a single function
		// print when "used" is false
		for i := 0; i < len_USD; i++ {
			date_range := con.Units.USD[i].Start + " - " + con.Units.USD[i].End
			if !used[date_range] {
				used[date_range] = true
				printUnit(con.Units.USD[i], &rowIter)
			}
		}

		/*
			// print out the frame
			for i := 0; i < len_USD; i++ {
				axis = getCol(i, rowIter)
				f.SetCellValue("Sheet1", axis, con.Units.USD[i].Frame)
			}

			// go to the next row
			rowIter++
			axis = incrRow(rowIter)

			// print out the date range
			for i := 0; i < len_USD; i++ {
				axis = getCol(i, rowIter)
				date_range := con.Units.USD[i].Start + " - " + con.Units.USD[i].End
				if strings.Compare(con.Units.USD[i].Frame, "") != 0 {
					f.SetCellValue("Sheet1", axis, date_range)
				}
			}

			// go to the next row
			rowIter++
			axis = incrRow(rowIter)

			// print out filing date
			for i := 0; i < len_USD; i++ {
				axis = getCol(i, rowIter)
				if strings.Compare(con.Units.USD[i].Frame, "") != 0 {
					f.SetCellValue("Sheet1", axis, con.Units.USD[i].Filed)
				}
			}

			// go to the next row
			rowIter++
			axis = incrRow(rowIter)

			// print out the values
			for i := 0; i < len_USD; i++ {
				axis = getCol(i, rowIter)
				if strings.Compare(con.Units.USD[i].Frame, "") != 0 {
					f.SetCellValue("Sheet1", axis, con.Units.USD[i].Val)
				}
			}
		*/
	} else if len(con.Units.EUR) > 0 {

	} else if len(con.Units.BRL) > 0 {

	} else if len(con.Units.Acre) > 0 {

	} else if len(con.Units.Shares) > 0 {

	} else {
		fmt.Println("unsupported unit type")
	}

	if err = f.SaveAs(filename); err != nil {
		return -1, err
	}

	rowIter++
	return rowIter, nil
}

// build a report with multiple concepts
func BuildCombinedReport(ticker string, facts *CompanyFacts, concepts []string) error {
	var filename = ticker + "_"
	for _, concept := range concepts {
		filename = filename + concept
	}
	filename = filename + "Report.xlsx"

	excelFile := excelize.NewFile()
	rowIter := 0

	for _, concept := range concepts {
		con, err := GetConcept(facts, concept)
		if err != nil {
			return err
		}
		rowIter, err = printConcept(con, excelFile, filename, rowIter)
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// build a report
func BuildReport(facts *CompanyFacts, concept string) error {
	con, err := GetConcept(facts, concept)
	if err != nil {
		return err
	}

	filename := concept + "Report.xlsx"
	excelFile := excelize.NewFile()
	_, err = printConcept(con, excelFile, filename, 0)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// returns a string of the CIK in the form used in API requests
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

// returns map of ticker --> CIK of all publicly traded companies in SEC's database
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

	ticker_to_CIK := &TickerToCIK{make(map[string]string)}
	for _, v := range coll.Collection {
		ticker_to_CIK.Collection[v.Ticker] = CIKToString(v.CIK)
	}

	return ticker_to_CIK, nil
}
