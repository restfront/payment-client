package payment

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// GetStatus запрашивает состояние банковского терминала или конкретной транзакции при непустом transactionID
func (t *bankTerminal) GetStatus(ctx context.Context, transactionID string) (*BankTerminalResponse, error) {
	path := "bank/status"

	if transactionID != "" {
		params := url.Values{}
		params.Add("transaction", transactionID)
		path = fmt.Sprintf("%s?%s", path, params.Encode())
	}

	result := &BankTerminalResponse{}

	_, err := t.doRequest(ctx, http.MethodGet, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе состояния банковского терминала: %w", err)
	}

	return result, nil
}

// TestHost проверяет соединение с банком
func (t *bankTerminal) TestHost(ctx context.Context) (*BankTerminalResponse, error) {
	path := "bank/test/host"

	result := &BankTerminalResponse{}

	_, err := t.doRequest(ctx, http.MethodGet, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при проверке соединения с банком: %w", err)
	}

	return result, nil
}

// TestPinpad проверяет соединение с пинпадом
func (t *bankTerminal) TestPinpad(ctx context.Context) (*BankTerminalResponse, error) {
	path := "bank/test/pinpad"

	result := &BankTerminalResponse{}

	_, err := t.doRequest(ctx, http.MethodGet, path, nil, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при проверке соединения с пинпадом: %w", err)
	}

	return result, nil
}

// InitiatePayment создает платеж
func (t *bankTerminal) InitiatePayment(ctx context.Context, payment BankPayment) (*BankTerminalResponse, error) {
	path := "bank/action/pay"

	result := &BankTerminalResponse{}

	_, err := t.doRequest(ctx, http.MethodPost, path, payment, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании платежа: %w", err)
	}

	return result, nil
}

// SubmitAction подтверждает действие
func (t *bankTerminal) SubmitAction(ctx context.Context, action BankTransactionAction) (*BankTerminalResponse, error) {
	path := "bank/answer"

	result := &BankTerminalResponse{}

	_, err := t.doRequest(ctx, http.MethodPost, path, action, result)
	if err != nil {
		return nil, fmt.Errorf("ошибка при подтверждении действия: %w", err)
	}

	return result, nil
}

// doRequest выполняет запрос к банковскому терминалу используя указанный метод, путь и тело запроса
func (t *bankTerminal) doRequest(
	ctx context.Context,
	method string,
	path string,
	body any,
	result any,
) (*resty.Response, error) {
	// формирование URL
	endpoint, err := url.JoinPath(t.parent.config.BaseURL, path)
	if err != nil {
		return nil, ErrIncorrectURL
	}

	// инициализация запроса
	req := t.parent.httpClient.R().
		SetContext(ctx).
		SetResult(result)

	// заголовок и тело запроса
	if body != nil {
		req.SetHeader("Content-Type", "application/json").
			SetBody(body)
	}

	// выполнение запроса
	response := &resty.Response{}

	switch method {
	case http.MethodGet:
		response, err = req.Get(endpoint)
	case http.MethodPost:
		response, err = req.Post(endpoint)
	default:
		return nil, fmt.Errorf("%w: %s", ErrIncorrectRequestMethod, method)
	}

	// обработка ошибок
	if err != nil {
		if isTimeout(err) {
			return nil, ErrConnectionTimeout
		}
		return nil, fmt.Errorf("ошибка при выполнении запроса: %w", err)
	}

	if response.IsError() {
		return nil, fmt.Errorf("ошибка при выполнении запроса (status: %d)", response.StatusCode())
	}

	return response, nil
}

func (f *fiscalRegister) OpenShift(ctx context.Context) error {

	return nil
}

func (f *fiscalRegister) CloseShift(ctx context.Context) error {

	return nil
}
