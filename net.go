package payment

import (
	"context"
	"errors"
	"net"
	"net/url"
	"os"
	"runtime"
	"strings"
	"syscall"
)

// isTimeout возвращает true для сетевых таймаутов.
func isTimeout(err error) bool {
	if err == nil {
		return false
	}

	// Любая ошибка, удовлетворяющая net.Error и Timeout()
	var ne net.Error
	if errors.As(err, &ne) && ne.Timeout() {
		return true
	}

	// Дедлайны
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, os.ErrDeadlineExceeded) {
		return true
	}

	// Обёртки уровня url/net/os.syscall
	var se *os.SyscallError
	if errors.As(err, &se) {
		// Unix: ETIMEDOUT
		if errors.Is(se.Err, syscall.ETIMEDOUT) {
			return true
		}
		// Windows: WSAETIMEDOUT (10060), без зависимости от x/sys/windows
		if errno, ok := se.Err.(syscall.Errno); ok && runtime.GOOS == "windows" && uint32(errno) == 10060 {
			return true
		}
	}

	// Последний безопасный fallback по известным сообщениям
	var ue *url.Error
	if errors.As(err, &ue) && strings.Contains(ue.Error(), "Client.Timeout exceeded") {
		return true
	}
	if strings.Contains(err.Error(), "i/o timeout") {
		return true
	}

	return false
}
