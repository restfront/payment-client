package payment

// // StartFiscalRegister - инициализация фискального регистратора
// func (client *Client) StartFiscalRegister() (*FiscalRegisterResponse, error) {
// 	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL)
// 	if err != nil {
// 		return nil, wrapURLError(err)
// 	}

// 	result := &FiscalRegisterResponse{}
// 	resultError := &errorResponse{}

// 	response, err := client.httpClient.R().
// 		SetHeader("Content-Type", "application/json").
// 		SetQueryParam("username", client.config.Username).
// 		SetResult(&result).
// 		SetError(&resultError).
// 		Get(endpoint)
// 	if err != nil {
// 		return nil, wrapConnectionError(err)
// 	}

// 	if response.IsError() {
// 		return nil, wrapRequestError(resultError.Error)
// 	}

// 	return result, nil
// }

// // OpenShift - открытие смены
// func (client *Client) OpenShift(request FiscalRegisterRequest) error {
// 	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL, "sessionStart")
// 	if err != nil {
// 		return wrapURLError(err)
// 	}

// 	resultError := &errorResponse{}

// 	response, err := client.httpClient.R().
// 		SetHeader("Content-Type", "application/json").
// 		SetBody(request).
// 		SetError(&resultError).
// 		Post(endpoint)
// 	if err != nil {
// 		return wrapConnectionError(err)
// 	}

// 	if response.IsError() {
// 		return wrapRequestError(resultError.Error)
// 	}

// 	return nil
// }

// // PrintSalesReceipt печать чека когда произведена продажа
// func (client *Client) PrintSalesReceipt(request ReceiptFiscalRegisterRequest) (*ReceiptFiscalRegisterResponse, error) {
// 	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL, "sell")
// 	if err != nil {
// 		return nil, wrapURLError(err)
// 	}

// 	result := &ReceiptFiscalRegisterResponse{}
// 	resultError := &errorResponse{}

// 	response, err := client.httpClient.R().
// 		SetHeader("Content-Type", "application/json").
// 		SetBody(request).
// 		SetResult(&result).
// 		SetError(&resultError).
// 		Post(endpoint)
// 	if err != nil {
// 		return nil, wrapConnectionError(err)
// 	}

// 	if response.IsError() {
// 		return nil, wrapRequestError(resultError.Error)
// 	}

// 	return result, nil
// }

// // CreateXReport печать X-отчета
// func (client *Client) CreateXReport(request FiscalRegisterRequest) (*ReportFiscalRegisterResponse, error) {
// 	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL, "reportX")
// 	if err != nil {
// 		return nil, wrapURLError(err)
// 	}

// 	result := &ReportFiscalRegisterResponse{}
// 	resultError := &errorResponse{}

// 	response, err := client.httpClient.R().
// 		SetHeader("Content-Type", "application/json").
// 		SetBody(request).
// 		SetResult(&result).
// 		SetError(&resultError).
// 		Post(endpoint)
// 	if err != nil {
// 		return nil, wrapConnectionError(err)
// 	}

// 	if response.IsError() {
// 		return nil, wrapRequestError(resultError.Error)
// 	}

// 	return result, nil
// }

// // CreateZReport печать Z-отчета
// func (client *Client) CreateZReport(request FiscalRegisterRequest) (*ReportFiscalRegisterResponse, error) {
// 	endpoint, err := url.JoinPath(client.config.BaseURL, registerURL, "reportZ")
// 	if err != nil {
// 		return nil, wrapURLError(err)
// 	}

// 	result := &ReportFiscalRegisterResponse{}
// 	resultError := &errorResponse{}

// 	response, err := client.httpClient.R().
// 		SetHeader("Content-Type", "application/json").
// 		SetBody(request).
// 		SetResult(&result).
// 		SetError(&resultError).
// 		Post(endpoint)
// 	if err != nil {
// 		return nil, wrapConnectionError(err)
// 	}

// 	if response.IsError() {
// 		return nil, wrapRequestError(resultError.Error)
// 	}

// 	return result, nil
// }
