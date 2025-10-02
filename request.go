package payment

type FiscalRegisterRequest struct {
	UserName string `json:"username"`
}

type ReceiptFiscalRegisterRequest struct {
	ID                 int    `json:"id"`
	UUID               string `json:"uuid"`
	IKassaSaleID       int    `json:"iKassaSaleId"`
	Number             string `json:"number"`
	SystemNumber       int    `json:"systemNumber"`
	ReturnSystemNumber int    `json:"returnSystemNumber"`
	NumberMPP          int    `json:"numberMPP"`
	UserName           string `json:"userName"`
	UserKey            int    `json:"userKey"`
	WaiterName         string `json:"WaiterName"`
	WaiterKey          int    `json:"WaiterKey"`
	DocumentID         string `json:"documentID"`
	ShiftNumber        int    `json:"ShiftNumber"`
	PrinterNoPrint     int    `json:"PrinterNoPrint"`
	CheckType          string `json:"checkType"`

	TotalSums TotalSums `json:"sums"`
	Lines     []Line    `json:"lines"`
	Payments  []Payment `json:"payments"`
	Comment   string    `json:"comment"`
}

// TotalSums итоговые суммы
type TotalSums struct {
	Amount       float64 `json:"allSum"`
	Cash         float64 `json:"cashSum"`
	Card         float64 `json:"cardSum"`
	Credit       float64 `json:"creditSum"`
	PersonalCard float64 `json:"personalCardSum"`
	Change       float64 `json:"changeSum"`
}

// Line позиция заказа
type Line struct {
	LineKey           int     `json:"lineKey"`
	GoodKey           int     `json:"goodKey"`
	GoodBarcodeGenKey int     `json:"goodBarcodeGenKey"`
	ScannedData       string  `json:"scannedData"`
	GoodBarcode       string  `json:"goodBarcode"`
	GoodDepartment    int     `json:"goodDepartment"`
	Quantity          float64 `json:"quantity"`
	Type              string  `json:"type"`
	GoodName          string  `json:"goodName"`
	CostNcu           float64 `json:"costNcu"`
	SumNcu            float64 `json:"sumNcu"`
	SumNcuDiscount    float64 `json:"sumNcuDiscount"`
	DiscountPercent   int     `json:"discountPercent"`
}

type Payment struct {
	SumNcu       float64 `json:"sumNcu"`
	Alias        string  `json:"alias"`
	PayName      string  `json:"payName"`
	PayGroupKey  int     `json:"PayGroupKey"`
	PayGroupName string  `json:"PayGroupName"`
	PayMethod    int     `json:"payMethod"`
}
