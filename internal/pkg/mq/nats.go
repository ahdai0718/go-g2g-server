package mq

import (
	"ohdada/g2gserver/internal/pkg/glog"

	"github.com/nats-io/nats.go"
)

type natsMessager struct {
	conn *nats.Conn
	jsc  nats.JetStreamContext
}

func (messager *natsMessager) Init(uris string) (err error) {

	messager.conn, err = nats.Connect(uris)
	if err != nil {
		glog.Error(err)
		return err
	}

	return
}

func (messager *natsMessager) Run() (err error) {
	return
}

func (messager *natsMessager) AddChannel(name string, subjects ...string) (err error) {

	return
}

func (messager *natsMessager) Close() (err error) {

	messager.conn.Close()
	err = messager.conn.Drain()

	return
}

func (messager *natsMessager) Subscribe(subject string, callback SubscribeCallbackFun) (subscription Subscription, err error) {

	subscription, err = messager.conn.Subscribe(subject, func(m *nats.Msg) {
		message := &Message{
			Subject: m.Subject,
			Reply:   m.Reply,
			Data:    m.Data,
		}
		callback(message)
	})

	if err != nil {
		glog.Error(err)
	}

	return
}

func (messager *natsMessager) Publish(subject string, message *Message) (err error) {

	err = messager.conn.Publish(subject, message.Data)

	if err != nil {
		glog.Error(err)
	}

	return
}
