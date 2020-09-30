package notifications

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"

	"github.com/spf13/viper"
)

// Email sends a notification to the designated sender via email
func email(subj string, body string) {
	from := mail.Address{"", viper.GetString("EMAIL_SENDER_FROM")}
	to := mail.Address{"", viper.GetString("EMAIL_TO")}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := viper.GetString("EMAIL_SENDER_HOST")

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", viper.GetString("EMAIL_SENDER_FROM"), viper.GetString("EMAIL_SENDER_PASSWORD"), host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(servername)
	if err != nil {
		log.Panic(err)
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

}
