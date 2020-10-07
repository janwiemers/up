package monitors

import (
	"net"
	"time"
)

func tcpMonitor(url string, expectation string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", url, timeout)
	if err != nil {
		return false
	}
	if conn != nil {
		defer conn.Close()
		return true
	}

	return false
}
