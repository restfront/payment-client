package payment

import "errors"

type TerminalOperationStatus string
type TransactionAction string
type ReceiptType string
type PaymentType string

const (
	TerminalOperationStatusSuccess    TerminalOperationStatus = "End"        // операция завершена успешно
	TerminalOperationStatusProcess    TerminalOperationStatus = "Process"    // операция в процессе выполнения
	TerminalOperationStatusFeedback   TerminalOperationStatus = "Feedback"   // ожидается ответное действие кассира/покупателя
	TerminalOperationStatusCancel     TerminalOperationStatus = "Cancel"     // операция отменена
	TerminalOperationStatusError      TerminalOperationStatus = "Error"      // ошибка при выполнении операции
	TerminalOperationStatusBusy       TerminalOperationStatus = "Occupied"   // терминал занят выполнением служебной или иной банковской операции
	TerminalOperationStatusIdle       TerminalOperationStatus = "NoProc"     // терминал находится в режиме простоя
	TerminalOperationStatusNextNumber TerminalOperationStatus = "NextNumber" // требование следующего номера операции
	TerminalOperationStatusUnknown    TerminalOperationStatus = "Unknown"    // результат операции неизвестен

	TransactionActionConfirm TransactionAction = "confirm"
	TransactionActionCancel  TransactionAction = "cancel"

	ReceiptTypeSell ReceiptType = "sell"

	PaymentTypeCash   PaymentType = "Cash"
	PaymentTypeCard   PaymentType = "Credit"
	PaymentTypeOplati PaymentType = "Oplati"
)

var (
	ErrIncorrectURL           = errors.New("некорректный URL")
	ErrConnectionTimeout      = errors.New("таймаут соединения/запроса")
	ErrIncorrectRequestMethod = errors.New("неподдерживаемый метод запроса")
	// ошибки, связанные с результатом платежа:
	ErrPaymentCanceled      = errors.New("платеж отменен пользователем/терминалом")
	ErrPaymentFailed        = errors.New("ошибка при осуществлении платежа")
	ErrPaymentUnknownStatus = errors.New("результат оплаты неизвестен")
	// ошибки, связанные с неожиданным состоянием терминала:
	ErrTerminalIdleUnexpected = errors.New("оплата завершена с неожиданным статусом")
	ErrTerminalNextNumber     = errors.New("терминал требует следующий номер операции")

	ErrFiscalRegisterShiftNotOpened     = errors.New("смена не открыта")
	ErrFiscalRegisterShiftLimitExceeded = errors.New("превышена максимальная продолжительность смены")
)

type BankTransactionAction struct {
	TransactionID int64             `json:"transaction"`
	Action        TransactionAction `json:"action"`
	Pin           string            `json:"pin,omitempty"`
}

type BankPayment struct {
	TransactionID int64   `json:"transaction"`
	Amount        float64 `json:"sum"`
	Currency      string  `json:"currency,omitempty"`
}

type BankTerminalResponse struct {
	Status        TerminalOperationStatus `json:"status"`
	Message       string                  `json:"message"`
	TransactionID *int64                  `json:"transaction"`
	AuthCode      *string                 `json:"authCode"`
	CardNumber    *string                 `json:"cardNumber"`
}

type BankTransaction struct {
	ID int64 `json:"transaction"`
}

type FiscalRegisterStatus struct {
	ShiftOpened        bool `json:"shiftOpen"`  // признак открытия смены
	ShiftLimitExceeded bool `json:"shiftEnded"` // признак превышения 24 часов с момента открытия смены
}

type FiscalRegisterShiftNumber struct {
	ShiftNumber int64 `json:"shiftnumber"` // номер смены
}

type FiscalRegisterPayment struct {
	ID           int64
	Number       string
	Cashier      string
	ReceiptType  ReceiptType
	Items        []FiscalRegisterPaymentItem
	PaymentTypes []FiscalRegisterPaymentType
}

type FiscalRegisterPaymentItem struct {
	ID          int64
	ProductID   int64
	ProductName string
	Quantity    int64
	Price       float64
	Amount      float64
}

type FiscalRegisterPaymentType struct {
	Type   PaymentType
	Amount float64
}

type FiscalRegisterPaymentResponse struct {
	DocumentID    string `json:"documentId"`
	ReceiptNumber int64  `json:"number"`
	ShiftNumber   int64  `json:"shiftNumber"`
}

type fiscalRegisterPaymentRequest struct {
	ID          int64  `json:"id"`
	Number      string `json:"number"`
	Cashier     string `json:"userName"`
	ReceiptType string `json:"checkType"`
	Totals      struct {
		CashAmount  float64 `json:"cashSum"`
		CardAmount  float64 `json:"cardSum"`
		OtherAmount float64 `json:"creditSum"`
		TotalAmount float64 `json:"allSum"`
	}
	Items        []fiscalRegisterPaymentItemRequest `json:"lines"`
	PaymentTypes []fiscalRegisterPaymentTypeRequest `json:"payments"`
}

type fiscalRegisterPaymentItemRequest struct {
	ID              int64   `json:"lineKey"`
	ProductID       int64   `json:"goodKey"`
	ProductName     string  `json:"goodName"`
	ProductType     string  `json:"type"`
	Quantity        int64   `json:"quantity"`
	Price           float64 `json:"costNcu"`
	Amount          float64 `json:"sumNcu"`
	DiscountPercent float64 `json:"discountPercent"`
	DiscountAmount  float64 `json:"sumNcuDiscount"`
	DepartmentID    int64   `json:"goodDepartment"`
}

type fiscalRegisterPaymentTypeRequest struct {
	Type   string  `json:"alias"`
	Amount float64 `json:"sumNcu"`
}

type fiscalRegisterOpenShiftRequest struct {
	Username string `json:"username"`
}

type fiscalRegisterCloseShiftRequest struct {
	Username string `json:"username"`
}

type fiscalRegisterPrintXReportRequest struct {
	Username string `json:"username"`
}

type fiscalRegisterPrintZReportRequest struct {
	Username string `json:"username"`
}

func (p *FiscalRegisterPayment) ToRequest() *fiscalRegisterPaymentRequest {
	request := &fiscalRegisterPaymentRequest{
		ID:          p.ID,
		Number:      p.Number,
		Cashier:     p.Cashier,
		ReceiptType: string(p.ReceiptType),
	}

	request.Items = make([]fiscalRegisterPaymentItemRequest, 0, len(p.Items))
	for _, item := range p.Items {
		request.Items = append(request.Items, *item.ToRequest())
	}

	request.PaymentTypes = make([]fiscalRegisterPaymentTypeRequest, 0, len(p.PaymentTypes))
	for _, paymentType := range p.PaymentTypes {
		request.PaymentTypes = append(request.PaymentTypes, *paymentType.ToRequest())
	}

	request.Totals.CashAmount = p.TotalCashAmount()
	request.Totals.CardAmount = p.TotalCardAmount()
	request.Totals.OtherAmount = p.TotalOtherAmount()
	request.Totals.TotalAmount = p.TotalAmount()

	return request
}

func (p FiscalRegisterPayment) TotalCashAmount() float64 {
	var total float64
	for _, paymentType := range p.PaymentTypes {
		if paymentType.Type == PaymentTypeCash {
			total += paymentType.Amount
		}
	}
	return total
}

func (p FiscalRegisterPayment) TotalCardAmount() float64 {
	var total float64
	for _, paymentType := range p.PaymentTypes {
		if paymentType.Type == PaymentTypeCard {
			total += paymentType.Amount
		}
	}
	return total
}

func (p FiscalRegisterPayment) TotalOtherAmount() float64 {
	var total float64
	for _, paymentType := range p.PaymentTypes {
		if paymentType.Type == PaymentTypeOplati {
			total += paymentType.Amount
		}
	}
	return total
}

func (p FiscalRegisterPayment) TotalAmount() float64 {
	var total float64
	for _, paymentType := range p.PaymentTypes {
		total += paymentType.Amount
	}
	return total
}

func (i *FiscalRegisterPaymentItem) ToRequest() *fiscalRegisterPaymentItemRequest {
	return &fiscalRegisterPaymentItemRequest{
		ID:              i.ID,
		ProductID:       i.ProductID,
		ProductName:     i.ProductName,
		ProductType:     "good",
		Quantity:        i.Quantity,
		Price:           i.Price,
		Amount:          i.Amount,
		DiscountPercent: 0,
		DiscountAmount:  0,
		DepartmentID:    0,
	}
}
func (t *FiscalRegisterPaymentType) ToRequest() *fiscalRegisterPaymentTypeRequest {
	return &fiscalRegisterPaymentTypeRequest{
		Type:   string(t.Type),
		Amount: t.Amount,
	}
}
