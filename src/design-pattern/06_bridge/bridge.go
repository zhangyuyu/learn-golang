package bridge

import "fmt"

// IMsgSender 发送消息接口
type IMsgSender interface {
	Send(msg string) string
}

type EmailMsgSender struct {
}

func (e *EmailMsgSender) Send(msg string) string {
	return fmt.Sprintf("Email: %s", msg)
}

// INotifaction 通知接口
type INotifaction interface {
	Notify(msg string) string
}

// ErrorNotification 错误通知
// 后面可能有warning 各种级别
type ErrorNotification struct {
	Sender IMsgSender
}

func (e *ErrorNotification) Notify(msg string) string {
	return e.Sender.Send(msg)
}
