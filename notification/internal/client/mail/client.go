package mail

import (
	"html/template"

	"github.com/wneessen/go-mail"
)

type Client interface {
	Send(to, subject string, bodyType mail.ContentType, body string) error
	SendHTMLTpl(to string, subject string, tpl *template.Template, data interface{}) error
	Close() error
}
