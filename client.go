// Package payment реализует клиент для работы с сервисом payment
package payment

import (
	"time"

	"github.com/go-resty/resty/v2"
)

var (
	defaultRequestTimeout = time.Second * 15
	defaultMaxAttempts    = 3
	defaultDelaySec       = 1
)

type Client struct {
	config     *Config
	httpClient *resty.Client
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
	httpClient.SetHeader("User-Agent", "RestFront/") // TODO name

	if config.DebugMode {
		httpClient.SetDebug(true)
	}

	client.httpClient = httpClient

	return client
}
