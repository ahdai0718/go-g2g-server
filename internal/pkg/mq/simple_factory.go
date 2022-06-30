package mq

import "ohdada/g2gserver/internal/pkg/glog"

const (
	Nats          = "nats"
	NatsJetStream = "nats_jet_stream"
)

var (
	simpleFactory = &SimpleFactory{}
)

func DefaultSimpleFactory() *SimpleFactory {
	return simpleFactory
}

type SimpleFactory struct{}

func (factory *SimpleFactory) Create(t string) (messager Messager) {

	switch t {
	case Nats:
		messager = &natsMessager{}
	case NatsJetStream:
		messager = &natsJetStreamMessager{}
	default:
		glog.Warningf("not match mq type [%s], failback to default.", t)
		messager = &natsMessager{}
	}

	return
}
