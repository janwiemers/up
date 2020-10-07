package monitors

import (
	"log"
	"net"
	"time"
)

func tcpMonitor(url string, expectation string) bool {
	timeout := time.Second
	_, err := net.DialTimeout("tcp", url, timeout)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
