package payment

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetStatus запрашивает состояние банковского терминала или конкретной транзакции при непустом transactionID
func (t *bankTerminal) GetStatus(ctx context.Context, transactionID int64) (*BankTerminalResponse, error) {
	path := "bank/status"

	params := url.Values{}
	if transactionID > 0 {
		params.Add("transaction", strconv.FormatInt(transactionID, 10))
	}

	result := &BankTerminalResponse{}

	_, err := t.parent.doRequest(ctx, http.MethodGet, path, params, nil, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе состояния банковского терминала: %w", err)
	}

	return result, nil
}

// TestHost проверяет соединение с банком
func (t *bankTerminal) TestHost(ctx context.Context) (*BankTerminalResponse, error) {
	path := "bank/test/host"

	result := &BankTerminalResponse{}

	_, err := t.parent.doRequest(ctx, http.MethodPost, path, nil, nil, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при проверке соединения с банком: %w", err)
	}

	return result, nil
}

// TestPinpad проверяет соединение с пинпадом
func (t *bankTerminal) TestPinpad(ctx context.Context) (*BankTerminalResponse, error) {
	path := "bank/test/pinpad"

	result := &BankTerminalResponse{}

	_, err := t.parent.doRequest(ctx, http.MethodPost, path, nil, nil, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при проверке соединения с пинпадом: %w", err)
	}

	return result, nil
}

// InitiatePayment создает платеж
func (t *bankTerminal) InitiatePayment(ctx context.Context, payment BankPayment) (*BankTerminalResponse, error) {
	path := "bank/pay"

	result := &BankTerminalResponse{}

	_, err := t.parent.doRequest(ctx, http.MethodPost, path, nil, payment, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании платежа: %w", err)
	}

	return result, nil
}

// SubmitAction подтверждает действие
func (t *bankTerminal) SubmitAction(ctx context.Context, action BankTransactionAction) (*BankTerminalResponse, error) {
	path := "bank/answer"

	result := &BankTerminalResponse{}

	_, err := t.parent.doRequest(ctx, http.MethodPost, path, nil, action, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подтверждении действия: %w", err)
	}

	return result, nil
}

func (f *fiscalRegister) OpenShift(ctx context.Context) error {

	return nil
}

func (f *fiscalRegister) CloseShift(ctx context.Context) error {

	return nil
}
