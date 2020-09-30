package notifications

import (
	"fmt"
	"time"
)

// Builder creates a subject and a message
func builder(service string, degraded bool) (string, string) {
	if degraded == true {
		return fmt.Sprintf("😲| Service %v is degraded", service), fmt.Sprintf("Degraded since %v", time.Now())
	}
	return fmt.Sprintf("🙂 | Service %v is restored", service), fmt.Sprintf("Restored at %v", time.Now())
}
