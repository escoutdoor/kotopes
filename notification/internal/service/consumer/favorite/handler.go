package favorite

import (
	"context"
	"encoding/json"
	"fmt"
	ht "html/template"
	"os"

	"github.com/IBM/sarama"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/notification/internal/model"
)

const (
	_tplFilename = "pkg/mailtpl/favorite.html"
)

func (svc *service) FavoriteHandler(ctx context.Context, msg *sarama.ConsumerMessage) error {
	const op = "favorite_consumer_service.FavoriteHandler"

	var data model.FavoriteMessage
	err := json.Unmarshal(msg.Value, &data)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	owner, err := svc.userClient.GetByID(ctx, data.OwnerID)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	if _, err := os.Stat(_tplFilename); os.IsNotExist(err) {
		return errwrap.Wrap(op, fmt.Errorf("mail template not found"))
	}

	htmlTpl, err := ht.ParseFiles(_tplFilename)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	tpldata := map[string]string{
		"FirstName": owner.FirstName,
		"LastName":  owner.LastName,
		"UserID":    data.UserID,
	}
	err = svc.mailClient.SendHTMLTpl(
		owner.Email,
		"Pet in favorite list",
		htmlTpl,
		tpldata,
	)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
