package payment

import "errors"

type TerminalOperationStatus string
type TransactionAction string

const (
	TerminalOperationStatusSuccess    TerminalOperationStatus = "End"        // операция завершена успешно
	TerminalOperationStatusProcess    TerminalOperationStatus = "Process"    // операция в процессе выполнения
	TerminalOperationStatusFeedback   TerminalOperationStatus = "Feedback"   // ожидается ответное действие кассира/покупателя
	TerminalOperationStatusCancel     TerminalOperationStatus = "Cancel"     // операция отменена
	TerminalOperationStatusError      TerminalOperationStatus = "Error"      // ошибка при выполнении операции
	TerminalOperationStatusBusy       TerminalOperationStatus = "Occupied"   // терминал занят выполнением служебной или иной банковской операции
	TerminalOperationStatusIdle       TerminalOperationStatus = "Idle"       // терминал находится в режиме простоя
	TerminalOperationStatusNextNumber TerminalOperationStatus = "NextNumber" // требование следующего номера операции
	TerminalOperationStatusUnknown    TerminalOperationStatus = "Unknown"    // результат операции неизвестен

	TransactionActionConfirm TransactionAction = "confirm"
	TransactionActionCancel  TransactionAction = "cancel"
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

var (
	ErrIncorrectURL           = errors.New("некорректный URL")
	ErrConnectionTimeout      = errors.New("таймаут соединения/запроса")
	ErrIncorrectRequestMethod = errors.New("неподдерживаемый метод запроса")
)
