package monitors

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/miekg/dns"
)

var dnsPort string = ":8002"

func TestDNSWithSuccess(t *testing.T) {
	srv := serverMockDNS()
	defer srv.Shutdown()

	res := dnsMonitor(fmt.Sprintf("127.0.0.1%v", dnsPort), "test.service")

	if res == false {
		t.Error("expected", true, "got", res)
	}
}

func TestDNSWithNoSuccess(t *testing.T) {
	res := dnsMonitor(fmt.Sprintf("127.0.0.1%v", dnsPort), "test.service")

	if res == true {
		t.Error("expected", false, "got", res)
	}
}

func TestDNSWithNoResults(t *testing.T) {
	srv := serverMockDNS()
	defer srv.Shutdown()
	res := dnsMonitor(fmt.Sprintf("127.0.0.1%v", dnsPort), "testing.service")

	if res == true {
		t.Error("expected", false, "got", res)
	}
}

func serverMockDNS() *dns.Server {
	dns.HandleFunc("service.", handleDNSRequest)
	server := &dns.Server{Addr: dnsPort, Net: "udp"}
	fmt.Printf("Starting at %v\n", dnsPort)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Failed to start server: %s\n ", err.Error())
		}
	}()
	time.Sleep(1 * time.Second)
	// returning reference so caller can call Shutdown()
	return server
}

var records = map[string]string{
	"test.service.": "1.2.3.4",
}

func parseQuery(m *dns.Msg) {
	for _, q := range m.Question {
		switch q.Qtype {
		case dns.TypeA:
			log.Printf("Query for %s\n", q.Name)
			ip := records[q.Name]
			if ip != "" {
				rr, err := dns.NewRR(fmt.Sprintf("%s A %s", q.Name, ip))
				if err == nil {
					m.Answer = append(m.Answer, rr)
				}
			}
		}
	}
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m)
	}

	w.WriteMsg(m)
}
