// Package main generate fake emails sent via SMTP
package main

import (
	"bytes"
	"flag"
	"net/smtp"
	"time"
)

var host = flag.String("host", "localhost:25", "host:port of SMTP server")

// message represents an unencoded email message
type message struct {
	from, to, subject, body string
}

func main() {
	flag.Parse()
	m := &message{
		"from@email",
		"to@email",
		"test subject",
		"amazing content goes here",
	}
	b := new(bytes.Buffer)
	if err := encodeMessage(b, m); err != nil {
		println(err)
		return
	}
	if err := sendMessage(*host, m, b); err != nil {
		println(err)
		return
	}
}

func encodeMessage(b *bytes.Buffer, m *message) error {
	b.Reset()
	b.WriteString("From: ")
	b.WriteString(m.from)
	b.WriteString("\r\nTo: ")
	b.WriteString(m.to)
	b.WriteString("\r\nDate: ")
	b.WriteString(time.Now().Format(time.RFC1123Z))
	b.WriteString("\r\nSubject: ")
	b.WriteString(m.subject)
	b.WriteString("\r\n\r\n")
	b.WriteString(m.body)
	b.WriteString("\r\n")
	return nil
}

func sendMessage(host string, m *message, b *bytes.Buffer) error {
	to := []string{m.to}
	return smtp.SendMail(host, nil, m.from, to, b.Bytes())
}
