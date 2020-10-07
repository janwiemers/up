package monitors

import (
	"context"
	"net"
	"time"
)

func dnsMonitor(address string, expectation string) bool {
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: time.Second,
			}
			return d.DialContext(ctx, "udp", address)
		},
	}
	ip, err := r.LookupIPAddr(context.Background(), expectation)

	if err != nil {
		return false
	}

	if len(ip) <= 0 {
		return false
	}

	return true
}
