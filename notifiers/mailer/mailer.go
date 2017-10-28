package mailer

// Mailer is an object used to send email
//go:generate mockgen -destination implementations/mockmailer/mailer.go -package mockmailer github.com/Nivl/go-rest-tools/notifiers/mailer Mailer
type Mailer interface {
	// SendStackTrace emails the current stacktrace to the default FROM
	SendStackTrace(trace []byte, message string, context map[string]string) error

	// Send is used to send an email
	Send(msg *Message) error
}
