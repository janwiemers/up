package monitors

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"
)

var tcpPort string = ":8000"

func TestTCPMonitorWithSuccess(t *testing.T) {
	srv := serverMockTCP()
	defer srv.Close()

	res := tcpMonitor(fmt.Sprintf("localhost%v", tcpPort), "")

	if res == false {
		t.Error("expected", true, "got", res)
	}
}

func TestTCPMonitorWithNoSuccess(t *testing.T) {
	res := tcpMonitor(tcpPort, "")

	if res == true {
		t.Error("expected", true, "got", res)
	}
}

func serverMockTCP() *http.Server {
	srv := &http.Server{Addr: tcpPort}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ok\n")
	})

	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return srv
}
