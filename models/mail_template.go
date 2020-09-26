package models

import (
	"context"
	"errors"
	"github.com/momenteam/momentum/database"
	"github.com/momenteam/momentum/models/enums"
)

type MailTemplate struct {
	TemplateType enums.MailTemplateType `bson:"mailTemplateType" json:"mailTemplateType"`
	Template string `bson:"template" json:"template"`
}

func CreateMailTemplate(mailTemplate MailTemplate) (result MailTemplate, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("mail template create error")
		}
	}()

	_, err = database.MailTemplateCollection.InsertOne(context.Background(), mailTemplate)

	return mailTemplate, err
}
