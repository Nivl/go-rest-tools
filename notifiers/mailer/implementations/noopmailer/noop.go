package noopmailer

import (
	"fmt"

	"github.com/Nivl/go-rest-tools/notifiers/mailer"
)

// Makes sure Noop implements Mailer
var _ mailer.Mailer = (*Noop)(nil)

// Noop is a mailer that just print emails
type Noop struct {
}

// SendStackTrace emails the current stacktrace to the default FROM
func (m *Noop) SendStackTrace(trace []byte, message string, context map[string]string) error {
	fmt.Printf("%s,%#v\n%s", message, context, trace)
	return nil
}

// Send is used to send an email
func (m *Noop) Send(msg *mailer.Message) error {
	fmt.Printf("FROM: %s\nTO: %s\nSUBJECT: %s\n%s\n", msg.From, msg.To, msg.Subject, msg.Body)
	return nil
}
