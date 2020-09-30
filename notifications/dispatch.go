package notifications

import (
	"github.com/spf13/viper"
)

// Dispatch is taking care of sending a message across all available channels
func Dispatch(service string, degraded bool) {
	subject, body := builder(service, degraded)

	if viper.GetBool("NOTIFICATIONS_ENABLE_EMAIL") == true {
		email(subject, body)
	}
}
