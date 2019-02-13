package messaging_test

import (
	"bytes"
	"testing"

	"github.com/annoying-external-dependency/mailer"
	"github.com/annoying-external-dependency/messaging"
	"github.com/stretchr/testify/assert"
)

func TestOnMsgOutputsContent(t *testing.T) {
	var buffer bytes.Buffer
	email := &mailer.Email{To: "a@b.c", Subject: "Spam", Body: "Egg"}

	messaging.OnMsg(email, &buffer)

	output := buffer.String()
	assert.Contains(t, output, "To: a[at]b.c")
	assert.Contains(t, output, "Body: Egg")
	assert.Contains(t, output, "Subject: Spam")
}
