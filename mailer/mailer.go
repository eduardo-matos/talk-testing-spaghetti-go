package mailer

import (
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

type Email struct {
	Body    string `json:"body"`
	Subject string `json:"subject"`
	To      string `json:"to"`
}

func ParseMsg(rawMsg amqp.Delivery) (*Email, error) {
	email := &Email{}

	err := json.Unmarshal(rawMsg.Body, email)

	if err != nil {
		return email, err
	}

	if !isEmail(email.To) {
		return email, errors.Errorf(`Invalid e-mail: "%s"`, email.To)
	}

	return email, nil
}

func isEmail(emailAddress string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(emailAddress) {
		return false
	}

	return true
}
