package mailer_test

import (
	"testing"

	"github.com/annoying-external-dependency/mailer"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func TestParseEmailFromValidJson(t *testing.T) {
	msg := amqp.Delivery{Body: []byte(`{
		"to": "dummy@spam.com",
		"subject": "Spam",
		"body": "Fish"
	}`)}

	result, err := mailer.ParseMsg(msg)

	assert.Nil(t, err)
	assert.Equal(t, result.To, "dummy@spam.com")
	assert.Equal(t, result.Body, "Fish")
	assert.Equal(t, result.Subject, "Spam")
}

func TestParseEmailFromInvalidValidJson(t *testing.T) {
	msg := amqp.Delivery{Body: []byte("{")}

	result, err := mailer.ParseMsg(msg)

	assert.NotNil(t, err)
	assert.Equal(t, result.To, "")
	assert.Equal(t, result.Body, "")
	assert.Equal(t, result.Subject, "")
}

func TestParseInvalidEmailFromValidJson(t *testing.T) {
	msg := amqp.Delivery{Body: []byte(`{
		"to": "dummy",
		"subject": "Spam",
		"body": "Fish"
	}`)}

	_, err := mailer.ParseMsg(msg)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), `Invalid e-mail: "dummy"`)
}
