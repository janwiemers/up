package monitors

import (
	"github.com/go-ping/ping"
)

func icmpMonitor(url string, expectation string) bool {
	pinger, err := ping.NewPinger(url)
	if err != nil {
		return false
	}
	pinger.Count = 1
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketLoss > 0.0 {
		return false
	}

	return true
}
