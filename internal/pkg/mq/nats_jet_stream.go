package mq

import (
	"ohdada/g2gserver/internal/pkg/glog"

	"github.com/nats-io/nats.go"
)

type natsJetStreamMessager struct {
	conn *nats.Conn
	jsc  nats.JetStreamContext
}

func (messager *natsJetStreamMessager) Init(uris string) (err error) {

	messager.conn, err = nats.Connect(uris)
	if err != nil {
		glog.Error(err)
		return err
	}

	// Create JetStream Context
	messager.jsc, err = messager.conn.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		glog.Error(err)
		return err
	}

	return
}

func (messager *natsJetStreamMessager) Run() (err error) {
	return
}

func (messager *natsJetStreamMessager) AddChannel(name string, subjects ...string) (err error) {

	_, err = messager.jsc.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: subjects,
	})

	if err != nil {
		glog.Error(err)
	}

	return
}

func (messager *natsJetStreamMessager) Close() (err error) {

	messager.conn.Close()
	err = messager.conn.Drain()

	return
}

func (messager *natsJetStreamMessager) Subscribe(subject string, callback SubscribeCallbackFun) (subscription Subscription, err error) {

	subscription, err = messager.jsc.Subscribe(subject, func(m *nats.Msg) {
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

func (messager *natsJetStreamMessager) Publish(subject string, message *Message) (err error) {

	_, err = messager.jsc.PublishAsync(subject, message.Data)

	if err != nil {
		glog.Error(err)
	}

	return
}
