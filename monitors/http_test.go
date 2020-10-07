package monitors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var httpPort = "0.0.0.0:8001"

func TestHttpMonitorWithSuccess(t *testing.T) {
	srv := serverMockHTTP()
	defer srv.Close()

	res := httpMonitor(fmt.Sprintf("%v/success", srv.URL), 200)

	if res == false {
		t.Error("expected", true, "got", res)
	}
}

func TestHttpMonitorWithNoSuccess(t *testing.T) {
	srv := serverMockHTTP()
	defer srv.Close()

	res := httpMonitor(fmt.Sprintf("%v/fail", srv.URL), 200)

	if res == true {
		t.Error("expected", false, "got", res)
	}
}

func TestHttpMonitorWithError(t *testing.T) {
	srv := serverMockHTTP()
	defer srv.Close()

	res := httpMonitor(fmt.Sprintf("1%v/success", srv.URL), 200)

	if res == true {
		t.Error("expected", false, "got", res)
	}
}

func serverMockHTTP() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/success", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("mock server responding"))
	})

	srv := httptest.NewServer(handler)

	return srv
}
