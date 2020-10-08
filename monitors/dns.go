package monitors

import (
	"github.com/lixiangzhong/dnsutil"
)

func dnsMonitor(address string, expectation string) bool {
	var dig dnsutil.Dig
	dig.SetDNS(address)
	a, err := dig.A(expectation)
	if err != nil {
		return false
	}

	if len(a) <= 0 {
		return false
	}

	return true
}
