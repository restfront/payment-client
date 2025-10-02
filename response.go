package payment

type errorResponse struct {
	Error string `json:"error"`
}

type FiscalRegisterResponse struct {
	ShiftEnded                  bool `json:"shiftEnded"`
	NoPaper                     bool `json:"noPaper"`
	PaperIsEnding               bool `json:"paperIsEnding"`
	BufferIsEnding              bool `json:"bufferIsEnding"`
	ShiftOpen                   bool `json:"shiftOpen"`
	SknoUse                     bool `json:"sknoUse"`
	SknoActive                  bool `json:"sknoActive"`
	Restaurant3                 bool `json:"restaurant3"`
	SknoDocNoTranferredInMemory bool `json:"sknoDocNoTranferredInMemory"`
}

type ReceiptFiscalRegisterResponse struct {
	FiscalregisterNumber int    `json:"fiscalregisternumber"`
	DocumentID           string `json:"documentID"`
	Number               int    `json:"number"`
	ShiftNumber          int    `json:"ShiftNumber"`
	TransactionReturnID  int    `json:"TransactionReturnId"`
}

type ReportFiscalRegisterResponse struct {
	Rnm                      string       `json:"Rnm"`
	SessionNumber            int          `json:"SessionNumber"`
	LastDocNumber            int          `json:"LastDocNumber"`
	FirstDocNumber           int          `json:"FirstDocNumber"`
	OpenDate                 string       `json:"OpenDate"`
	DeviceID                 string       `json:"DeviceId"`
	AddThrRollbackCount      int          `json:"AddThrRollbackCount"`
	AddThrRollbackSum        float64      `json:"AddThrRollbackSum"`
	AddThrTotalCount         int          `json:"AddThrTotalCount"`
	AddThrTotalSum           float64      `json:"AddThrTotalSum"`
	AdditOrdersCanceledCount int          `json:"AdditOrdersCanceledCount"`
	AddOrdersCanceledSum     float64      `json:"AddOrdersCanceledSum"`
	AddOrdersClosedCount     int          `json:"AddOrdersClosedCount"`
	AddOrdersClosedSum       float64      `json:"AddOrdersClosedSum"`
	AddOrdersCorrectedCount  int          `json:"AddOrdersCorrectedCount"`
	AddOrdersCorrectedSum    float64      `json:"AddOrdersCorrectedSum"`
	AddOrdersMovedCount      int          `json:"AddOrdersMovedCount"`
	AddOrdersMovedSum        float64      `json:"AddOrdersMovedSum"`
	AddOrdersTotalCount      int          `json:"AddOrdersTotalCount"`
	AddOrdersTotalSum        float64      `json:"AddOrdersTotalSum"`
	ClientWithdrawsSum       float64      `json:"ClientWithdrawsSum"`
	ClientWithdrawsCount     int          `json:"ClientWithdrawsCount"`
	CorrectionsSum           float64      `json:"CorrectionsSum"`
	CorrectionsCount         int          `json:"CorrectionsCount"`
	MoneyInSum               float64      `json:"MoneyInSum"`
	MoneyInCount             int          `json:"MoneyInCount"`
	MoneyOutSum              float64      `json:"MoneyOutSum"`
	MoneyOutCount            int          `json:"MoneyOutCount"`
	SellSum                  float64      `json:"SellSum"`
	SellCount                int          `json:"sellCount"`
	SellOther                float64      `json:"SellOther"`
	SellCard                 float64      `json:"SellCard"`
	SellSISum                float64      `json:"SellSISum"`
	SellSICount              int          `json:"SellSICount"`
	SellUKZSum               float64      `json:"SellUKZSum"`
	SellUKZCount             int          `json:"SellUKZCount"`
	RefundsSum               float64      `json:"RefundsSum"`
	RefundsCount             int          `json:"RefundsCount"`
	RefundsCard              float64      `json:"RefundsCard"`
	RefundsCash              float64      `json:"RefundsCash"`
	RefundsOther             float64      `json:"RefundsOther"`
	RefundsSISum             float64      `json:"RefundsSISum"`
	RefundsSICount           int          `json:"RefundsSICount"`
	RefundsUKZSum            float64      `json:"RefundsUKZSum"`
	RefundsUKZCount          int          `json:"RefundsUKZCount"`
	CancellationSum          float64      `json:"CancelationSum"`
	CancellationCount        int          `json:"CancelationCount"`
	CancellationCard         float64      `json:"CancelationCard"`
	CancellationCash         float64      `json:"CancelationCash"`
	CancellationOther        float64      `json:"CancelationOther"`
	CancellationSISum        float64      `json:"CancelationSISum"`
	CancellationSICount      int          `json:"CancelationSICount"`
	CancellationUKZSum       float64      `json:"CancelationUKZSum"`
	CancellationUKZCount     int          `json:"CancelationUKZCount"`
	CancelSum                float64      `json:"CancelSum"`
	CancelCount              int          `json:"CancelCount"`
	CashSum                  float64      `json:"CashSum"`
	CashRefundSum            float64      `json:"CashRefundSum"`
	DiscountSum              float64      `json:"DiscountSum"`
	StoredDocuments          int          `json:"StoredDocuments"`
	CashInBox                float64      `json:"CashInBox"`
	FregType                 int          `json:"FregType"`
	Items                    []ReportItem `json:"items"`
}

type ReportItem struct {
	Index       int     `json:"index"`
	SellSum     float64 `json:"sellSum"`
	RefundSum   float64 `json:"refundSum"`
	Alias       string  `json:"alias"`
	FiscalName  string  `json:"fiscalName"`
	ServiceName string  `json:"serviceName"`
}
