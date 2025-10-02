// Package payment реализует клиент для работы с сервисом payment
package payment

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	defaultRequestTimeout = time.Second * 15
	defaultMaxAttempts    = 3
	defaultDelaySec       = 1
)

// BankTerminal определяет методы для работы с банковским терминалом
type BankTerminal interface {
	GetStatus(ctx context.Context, transactionID string) (*BankTerminalResponse, error)
	SubmitAction(ctx context.Context, action BankTransactionAction) (*BankTerminalResponse, error)
	InitiatePayment(ctx context.Context, payment BankPayment) (*BankTerminalResponse, error)
	TestHost(ctx context.Context) (*BankTerminalResponse, error)
	TestPinpad(ctx context.Context) (*BankTerminalResponse, error)
}

type FiscalRegister interface {
	OpenShift(ctx context.Context) error
	CloseShift(ctx context.Context) error
}

type Client struct {
	config     *Config
	httpClient *resty.Client

	bankTerminal   BankTerminal
	fiscalRegister FiscalRegister
}

type Config struct {
	BaseURL           string
	Username          string
	RequestTimeoutSec int
	RetryPolicy       struct {
		MaxAttempts int
		DelaySec    int
	}
	DebugMode bool
}

type bankTerminal struct {
	parent *Client
}
type fiscalRegister struct {
	parent *Client
}

func NewClient(config *Config) *Client {
	client := &Client{
		config: config,
	}

	httpClient := resty.New()

	timeout := defaultRequestTimeout
	if config.RequestTimeoutSec > 0 {
		timeout = time.Duration(config.RequestTimeoutSec) * time.Second
	}

	maxAttemps := defaultMaxAttempts
	if config.RetryPolicy.MaxAttempts > 0 {
		maxAttemps = config.RetryPolicy.MaxAttempts
	}

	delaySec := defaultDelaySec
	if config.RetryPolicy.DelaySec > 0 {
		delaySec = config.RetryPolicy.DelaySec
	}

	httpClient.SetTimeout(timeout)
	httpClient.SetRetryCount(maxAttemps)
	httpClient.SetRetryWaitTime(time.Duration(delaySec) * time.Second)
	httpClient.SetHeader("User-Agent", "RestFront/payment-client")

	if config.DebugMode {
		httpClient.SetDebug(true)
	}

	client.httpClient = httpClient

	client.bankTerminal = &bankTerminal{parent: client}
	client.fiscalRegister = &fiscalRegister{parent: client}

	return client
}
