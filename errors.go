package payment

import (
	"errors"
	"fmt"
)

var (
	ErrConnection = errors.New("ошибка соединения с сервером")
	ErrRequest    = errors.New("ошибка запроса")
	ErrURL        = errors.New("ошибка формирования URL")
)

func wrapConnectionError(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%w: %v", ErrConnection, err)
}

func wrapRequestError(err string) error {
	return fmt.Errorf("%w: %v", ErrRequest, err)
}

func wrapURLError(err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%w: %v", ErrURL, err)
}
