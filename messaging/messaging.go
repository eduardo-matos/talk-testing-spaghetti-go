package messaging

import (
	"fmt"
	"io"
	"strings"

	"github.com/annoying-external-dependency/mailer"
)

func OnMsg(email *mailer.Email, writer io.Writer) {
	writer.Write([]byte("--\n"))
	writer.Write([]byte(fmt.Sprintf("To: %s\n", strings.Replace(email.To, "@", "[at]", 1))))
	writer.Write([]byte(fmt.Sprintf("Subject: %s\n", email.Subject)))
	writer.Write([]byte(fmt.Sprintf("Body: %s\n", email.Body)))
	writer.Write([]byte("--\n"))
}
