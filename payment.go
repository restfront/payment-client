package payment

import (
	"fmt"
	"net/url"
)

const (
	bankURL     string = "bank"
	registerURL string = "register"
)

// GetBankOperationStatus -
func (client *Client) GetBankOperationStatus(transactionID *string) (*BankResponse, error) {
	// TODO параметр transactionID обязательный или нет?
	endpoint, err := url.JoinPath(client.config.BaseURL, bankURL, "status")
	if err != nil {
		return nil, wrapURLError(err)
	}

	result := &BankResponse{}
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetResult(&result).
		SetError(&resultError).
		Get(endpoint)
	if err != nil {
		return nil, wrapConnectionError(err)
	}

	if response.IsError() {
		// TODO какая будет ошибка
		return nil, fmt.Errorf("ошибка при запросе на получение статуса операции банка: %s", resultError.Error)
	}

	return result, nil
}

// CreateAnswerToBank -
func (client *Client) CreateAnswerToBank(request BankAnswerRequest) error {
	endpoint, err := url.JoinPath(client.config.BaseURL, bankURL, "answer")
	if err != nil {
		return wrapURLError(err)
	}

	var result any
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&result).
		SetError(&resultError).
		Post(endpoint)
	if err != nil {
		return wrapConnectionError(err)
	}

	if response.IsError() {
		return fmt.Errorf("ошибка при выполнении запроса на формирование ответа для банка: %s", resultError.Error)
	}

	return nil
}

// CreateBankPayment -
func (client *Client) CreateBankPayment(request BankActionPayRequest) (*BankActionPayResponse, error) {
	endpoint, err := url.JoinPath(client.config.BaseURL, bankURL, "action/pay")
	if err != nil {
		return nil, wrapURLError(err)
	}

	result := &BankActionPayResponse{}
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&result).
		SetError(&resultError).
		Post(endpoint)
	if err != nil {
		return nil, wrapConnectionError(err)
	}

	if response.IsError() {
		return nil, fmt.Errorf("ошибка при выполнении запроса на создание оплаты: %s", resultError.Error)
	}

	return result, nil
}

// CreateTestConnection -
func (client *Client) CreateTestConnection(request BankActionPayRequest) (*BankActionPayResponse, error) {
	// NOTE: проверить реализацию

	endpoint, err := url.JoinPath(client.config.BaseURL, bankURL, "test/host")
	if err != nil {
		return nil, wrapURLError(err)
	}

	result := &BankActionPayResponse{}
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&result).
		SetError(&resultError).
		Post(endpoint)
	if err != nil {
		return nil, wrapConnectionError(err)
	}

	if response.IsError() {
		return nil, fmt.Errorf("ошибка при проверке соединения: %s", resultError.Error)
	}

	return result, nil
}

// CreateTestPinpad -
func (client *Client) CreateTestPinpad(request BankActionPayRequest) (*BankActionPayResponse, error) {
	// NOTE: проверить реализацию

	endpoint, err := url.JoinPath(client.config.BaseURL, bankURL, "test/pinpad")
	if err != nil {
		return nil, wrapURLError(err)
	}

	result := &BankActionPayResponse{}
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&result).
		SetError(&resultError).
		Post(endpoint)
	if err != nil {
		return nil, wrapConnectionError(err)
	}

	if response.IsError() {
		return nil, fmt.Errorf("ошибка при проверке отправки пинпад: %s", resultError.Error)
	}

	return result, nil
}

// StartFiscalRegister - инициализация фискального регистратора
func (client *Client) StartFiscalRegister() (*FiscalRegisterResponse, error) {
	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL)
	if err != nil {
		return nil, wrapURLError(err)
	}

	result := &FiscalRegisterResponse{}
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("username", client.config.Username).
		SetResult(&result).
		SetError(&resultError).
		Get(endpoint)
	if err != nil {
		return nil, wrapConnectionError(err)
	}

	if response.IsError() {
		return nil, wrapRequestError(resultError.Error)
	}

	return result, nil
}

// OpenShift - открытие смены
func (client *Client) OpenShift(request FiscalRegisterRequest) error {
	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL, "sessionStart")
	if err != nil {
		return wrapURLError(err)
	}

	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetError(&resultError).
		Post(endpoint)
	if err != nil {
		return wrapConnectionError(err)
	}

	if response.IsError() {
		return wrapRequestError(resultError.Error)
	}

	return nil
}

// PrintSalesReceipt печать чека когда произведена продажа
func (client *Client) PrintSalesReceipt(request ReceiptFiscalRegisterRequest) (*ReceiptFiscalRegisterResponse, error) {
	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL, "sell")
	if err != nil {
		return nil, wrapURLError(err)
	}

	result := &ReceiptFiscalRegisterResponse{}
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&result).
		SetError(&resultError).
		Post(endpoint)
	if err != nil {
		return nil, wrapConnectionError(err)
	}

	if response.IsError() {
		return nil, wrapRequestError(resultError.Error)
	}

	return result, nil
}

// CreateXReport печать X-отчета
func (client *Client) CreateXReport(request FiscalRegisterRequest) (*ReportFiscalRegisterResponse, error) {
	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL, "reportX")
	if err != nil {
		return nil, wrapURLError(err)
	}

	result := &ReportFiscalRegisterResponse{}
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&result).
		SetError(&resultError).
		Post(endpoint)
	if err != nil {
		return nil, wrapConnectionError(err)
	}

	if response.IsError() {
		return nil, wrapRequestError(resultError.Error)
	}

	return result, nil
}

// CreateZReport печать Z-отчета
func (client *Client) CreateZReport(request FiscalRegisterRequest) (*ReportFiscalRegisterResponse, error) {
	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL, "reportZ")
	if err != nil {
		return nil, wrapURLError(err)
	}

	result := &ReportFiscalRegisterResponse{}
	resultError := &errorResponse{}

	response, err := client.httpClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&result).
		SetError(&resultError).
		Post(endpoint)
	if err != nil {
		return nil, wrapConnectionError(err)
	}

	if response.IsError() {
		return nil, wrapRequestError(resultError.Error)
	}

	return result, nil
}
