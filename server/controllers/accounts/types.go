package accounts

// BalanceSheet represents the structure of the balance sheet data.
type BalanceSheet struct {
	Reports []Report `json:"Reports"`
}

type Report struct {
	ReportID       string   `json:"ReportID"`
	ReportName     string   `json:"ReportName"`
	ReportType     string   `json:"ReportType"`
	ReportTitles   []string `json:"ReportTitles"`
	ReportDate     string   `json:"ReportDate"`
	UpdatedDateUTC string   `json:"UpdatedDateUTC"`
	Rows           []Row    `json:"Rows"`
}

type Row struct {
	RowType string `json:"RowType"`
	Title   string `json:"Title,omitempty"`
	Rows    []Row  `json:"Rows,omitempty"`
	Cells   []Cell `json:"Cells,omitempty"`
}

type Cell struct {
	Value      string      `json:"Value"`
	Attributes []Attribute `json:"Attributes,omitempty"`
}

type Attribute struct {
	Value string `json:"Value"`
	Id    string `json:"Id"`
}
