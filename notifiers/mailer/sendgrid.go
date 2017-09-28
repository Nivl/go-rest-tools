package mailer

import (
	"errors"
	"fmt"
	"strings"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// Makes sure Sendgrid implements Mailer
var _ Mailer = (*Sendgrid)(nil)

// Sendgrid is an object used to send email through sendgrid
type Sendgrid struct {
	APIKey               string
	DefaultFrom          string
	DefaultTo            string
	StacktraceTemplateID string
}

// SendStackTrace emails the current stacktrace to the default FROM
func (s *Sendgrid) SendStackTrace(trace []byte, message string, context map[string]string) error {
	if s.StacktraceTemplateID == "" {
		return errors.New("StacktraceTemplateID not set")
	}

	msg := NewMessage(s.StacktraceTemplateID)
	stacktrace := string(trace[:])

	msg.Body = strings.Replace(stacktrace, "\n", "<br>", -1)

	htmlContext := fmt.Sprintf("<li><strong>Error</strong>: %s</li>\n", message)
	for k, v := range context {
		htmlContext += fmt.Sprintf("<li><strong>%s</strong>: %s</li>\n", k, v)
	}
	msg.SetVar("context", htmlContext)
	return s.Send(msg)
}

// Send is used to send an email
func (s *Sendgrid) Send(msg *Message) error {
	from := mail.NewEmail("No Reply", msg.From)
	if msg.From == "" {
		from = mail.NewEmail("No Reply", s.DefaultFrom)
	}

	to := mail.NewEmail(msg.To, msg.To)
	if msg.From == "" {
		to = mail.NewEmail(s.DefaultTo, s.DefaultTo)
	}

	content := mail.NewContent("text/html", msg.Body)
	email := mail.NewV3MailInit(from, msg.Subject, to, content)
	email.SetTemplateID(msg.TemplateID)

	for k, v := range msg.Vars {
		email.Personalizations[0].SetSubstitution(k, v)
	}

	request := sendgrid.GetRequest(s.APIKey, "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	_, err := sendgrid.API(request)

	return err
}

// NewSendgrid creates and returns a new sendgrid instance
func NewSendgrid(APIKey, defaultFrom, defaultTo, stacktraceUUID string) *Sendgrid {
	return &Sendgrid{
		APIKey:               APIKey,
		DefaultFrom:          defaultFrom,
		DefaultTo:            defaultTo,
		StacktraceTemplateID: stacktraceUUID,
	}
}
