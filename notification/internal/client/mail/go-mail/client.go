package go_mail

import (
	"html/template"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	def "github.com/escoutdoor/kotopes/notification/internal/client/mail"
	"github.com/escoutdoor/kotopes/notification/internal/config"
	"github.com/wneessen/go-mail"
)

type client struct {
	mailClient *mail.Client
	cfg        config.MailConfig
}

var _ def.Client = (*client)(nil)

func NewClient(cfg config.MailConfig) (*client, error) {
	cl, err := mail.NewClient(cfg.Host(),
		mail.WithPort(cfg.Port()),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.From()),
		mail.WithPassword(cfg.Password()),
	)
	if err != nil {
		return nil, err
	}

	return &client{
		mailClient: cl,
		cfg:        cfg,
	}, nil
}

func (cl *client) Close() error {
	return cl.mailClient.Close()
}

func (cl *client) Send(to, subject string, bodyType mail.ContentType, body string) error {
	const op = "mail_client.Send"

	msg := mail.NewMsg()

	err := msg.From(cl.cfg.From())
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	err = msg.To(to)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	msg.Subject(subject)
	msg.SetBodyString(bodyType, body)

	err = cl.mailClient.DialAndSend(msg)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}

func (cl *client) SendHTMLTpl(to string, subject string, tpl *template.Template, data interface{}) error {
	const op = "mail_client.SendHTMLTpl"

	msg := mail.NewMsg()

	err := msg.From(cl.cfg.From())
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	err = msg.To(to)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	msg.Subject(subject)

	err = msg.SetBodyHTMLTemplate(tpl, data)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	err = cl.mailClient.DialAndSend(msg)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
