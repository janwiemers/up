package monitors

import (
	"testing"
)

func TestICMPMonitorWithSuccess(t *testing.T) {
	res := icmpMonitor("127.0.0.1", "")

	if res == false {
		t.Error("expected", true, "got", res)
	}
}

func TestICMPMonitorWithNoSuccess(t *testing.T) {
	res := icmpMonitor("bocalhost", "")

	if res == true {
		t.Error("expected", true, "got", res)
	}
}
