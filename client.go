// Package payment реализует клиент для работы с сервисом payment
package payment

import (
	"context"
	"fmt"
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
	GetStatus(ctx context.Context, transactionID int64) (*BankTerminalResponse, error)
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
	DevMode   bool
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

func (c *Client) CheckBankHost(ctx context.Context) error {
	// в режиме разработки не проверяем доступность соединения с банком
	if c.config.DevMode {
		return nil
	}

	resp, err := c.bankTerminal.TestHost(ctx)
	if err != nil {
		return err
	}

	if resp.Status != TerminalOperationStatusProcess {
		return fmt.Errorf("ошибка при проверке соединения с банком: %s", resp.Message)
	}

	attempts := 5
	delay := 1 * time.Second

	err = retry(ctx, attempts, delay, func(ctx context.Context) (bool, error) {
		resp, err := c.bankTerminal.GetStatus(ctx, 0)
		if err != nil {
			return false, err
		}
		return resp.Status == TerminalOperationStatusSuccess, nil
	})
	if err != nil {
		return fmt.Errorf("ошибка при проверке соединения с банком: %w", err)
	}

	return nil
}

func (c *Client) CheckBankPinpad(ctx context.Context) error {
	resp, err := c.bankTerminal.TestPinpad(ctx)
	if err != nil {
		return err
	}

	if resp.Status != TerminalOperationStatusProcess {
		return fmt.Errorf("ошибка при проверке соединения с пинпадом: %s", resp.Message)
	}

	attempts := 5
	delay := 1 * time.Second

	err = retry(ctx, attempts, delay, func(ctx context.Context) (bool, error) {
		resp, err := c.bankTerminal.GetStatus(ctx, 0)
		if err != nil {
			return false, err
		}
		return resp.Status == TerminalOperationStatusSuccess, nil
	})
	if err != nil {
		return fmt.Errorf("ошибка при проверке соединения с пинпадом: %w", err)
	}

	return nil
}

func (c *Client) ProcessBankPayment(ctx context.Context, payment BankPayment) (*BankTerminalResponse, error) {
	bankTerminal := c.bankTerminal

	// ожидаем готовность терминала
	attempts := 5
	delay := 1 * time.Second

	err := retry(ctx, attempts, delay, func(ctx context.Context) (bool, error) {
		resp, err := bankTerminal.GetStatus(ctx, 0)
		if err != nil {
			return false, err
		}
		return resp.Status == TerminalOperationStatusIdle, nil
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при ожидании готовности терминала: %w", err)
	}

	// проверяем доступность соединения с банком
	err = c.CheckBankHost(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при осуществлении платежа: %w", err)
	}

	// проверяем доступность соединения с пинпадом
	err = c.CheckBankPinpad(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при осуществлении платежа: %w", err)
	}

	// инициируем процес оплаты
	resp, err := bankTerminal.InitiatePayment(ctx, payment)
	if err != nil {
		return nil, fmt.Errorf("ошибка при осуществлении платежа: %w", err)
	}

	// ожидаем завершения оплаты
	attempts = 15
	delay = 1 * time.Second
	err = retry(ctx, attempts, delay, func(ctx context.Context) (bool, error) {
		resp, err := bankTerminal.GetStatus(ctx, payment.TransactionID)
		if err != nil {
			return false, err
		}

		switch resp.Status {
		case TerminalOperationStatusSuccess:
			// при успешном статусе завершаем цикл
			return true, nil
		case TerminalOperationStatusFeedback:
			// при обратном вызове повторяем цикл
			return false, nil
		case TerminalOperationStatusCancel:
			// при отмене платежа завершаем цикл
			return true, fmt.Errorf("отмена платежа: %s", resp.Message)
		case TerminalOperationStatusError:
			// при ошибке завершаем цикл
			return true, fmt.Errorf("ошибка при осуществлении платежа: %s", resp.Message)
		case TerminalOperationStatusBusy:
			// при занятости повторяем цикл
			return false, nil
		case TerminalOperationStatusProcess:
			// при ожидании завершения повторяем цикл
			return false, nil
		default:
			// при других статусах повторяем цикл
			return false, nil
		}
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при ожидании завершения платежа: %w", err)
	}

	return resp, nil
}

func retry(
	ctx context.Context,
	attempts int,
	delay time.Duration,
	fn func(context.Context) (bool, error),
) error {
	for i := 1; i <= attempts; i++ {
		ok, err := fn(ctx)
		if err != nil {
			return err
		}
		if ok {
			return nil
		}

		if i < attempts {
			select {
			case <-time.After(delay * time.Duration(i)):
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}

	return fmt.Errorf("достигнуто максимальное число попыток (%d)", attempts)
}
