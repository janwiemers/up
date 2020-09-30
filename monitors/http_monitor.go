package monitors

import (
	"log"
	"net/http"
)

func httpMonitor(url string, expectation int) bool {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return false
	}

	if resp.StatusCode != expectation {
		return false
	}

	return true
}
