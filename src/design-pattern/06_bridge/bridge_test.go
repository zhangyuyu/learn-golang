package bridge

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorNotification_Notify(t *testing.T) {
	sender := &EmailMsgSender{}
	notification := &ErrorNotification{
		Sender: sender,
	}
	result := notification.Notify("Hello, msg!")

	assert.Equal(t, "Email: Hello, msg!", result)
}
