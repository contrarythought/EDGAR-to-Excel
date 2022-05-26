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
			ret = append(ret, k)
		}
	} else if strings.Compare(accounting, "ifrs-full") == 0 {
		for k := range facts.Facts.IFRSFull {
			ret = append(ret, k)
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
func incrRow(axis string, index int) (string, error) {
	var ret string
	ret = "A" + strconv.Itoa(index)
	return ret, nil
}

// FIXME
func incrCol(axis string, index int) (string, error) {
	ret := ""

	letter_portion := axis[:len(axis)-1]
	old_letters := ""
	len_letter_portion := len(letter_portion)
	if len_letter_portion > 1 {
		old_letters = letter_portion[:len(letter_portion)-1]
	} else {
		old_letters = letter_portion
	}

	last_letter := ""
	new_letter := ""
	if index%26 == 0 {
		new_letter = "A"
	} else if len_letter_portion > 1 {
		last_letter = string(letter_portion[len(letter_portion)-1] + 1)
	} else {
		old_letters = string(old_letters[0] + 1)
	}

	ret = old_letters + last_letter + new_letter + string(axis[len(axis)-1])
	return ret, nil
}

// TODO - fix formatting
func printConcept(con *CompanyConcept, f *excelize.File, filename string) error {
	rowIter := 1
	axis := "A" + strconv.Itoa(rowIter)
	fmt.Println("Starting axis: ", axis)
	var err error
	if len(con.Units.USD) > 0 {
		// print out the form type
		for i := 0; i < len(con.Units.USD); i++ {
			f.SetCellValue("Sheet1", axis, con.Units.USD[i].Form)
			// increment the col
			axis, err = incrCol(axis, i+1)
			fmt.Println("Axis: ", axis)
			if err != nil {
				return err
			}
		}
		// go to the next row
		rowIter++
		axis, err = incrRow(axis, rowIter)
		if err != nil {
			return err
		}
		// print out the date range
		for i := 0; i < len(con.Units.USD); i++ {
			date_range := con.Units.USD[i].Start + " - " + con.Units.USD[i].End
			f.SetCellValue("Sheet1", axis, date_range)
			axis, err = incrCol(axis, i+1)
			if err != nil {
				return err
			}
			fmt.Println("Axis: ", axis)
		}
		// go to the next row
		rowIter++
		axis, err = incrRow(axis, rowIter)
		if err != nil {
			return err
		}
		// print out the values
		for i := 0; i < len(con.Units.USD); i++ {
			f.SetCellValue("Sheet1", axis, con.Units.USD[i].Val)
			axis, err = incrCol(axis, i+1)
			if err != nil {
				return err
			}
		}
	} else if len(con.Units.EUR) > 0 {

	} else if len(con.Units.BRL) > 0 {

	} else if len(con.Units.Acre) > 0 {

	} else if len(con.Units.Shares) > 0 {

	} else {
		fmt.Println("unsupported unit type")
	}

	if err = f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}

// build a report with multiple concepts
func BuildCombinedReport(facts *CompanyFacts, concepts []string) error {
	var filename string
	for _, concept := range concepts {
		filename = filename + concept
	}
	filename = filename + "Report.xlsx"
	/*
		f, err := os.Create(filename + "Report.txt")
		if err != nil {
			return err
		}
		defer f.Close()
	*/

	excelFile := excelize.NewFile()
	for _, concept := range concepts {
		con, err := GetConcept(facts, concept)
		if err != nil {
			return err
		}
		err = printConcept(con, excelFile, filename)
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
	/*
		f, err := os.Create(concept + "Report.txt")
		if err != nil {
			return err
		}
		defer f.Close()
	*/

	excelFile := excelize.NewFile()
	err = printConcept(con, excelFile, filename)
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
