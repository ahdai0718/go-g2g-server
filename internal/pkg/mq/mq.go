package mq

type Messager interface {
	Subscriber
	Publisher
	Init(uris string) error
	Run() error
	AddChannel(name string, subjects ...string) error
	Close() error
}

type Message struct {
	Subject string
	Reply   string
	Data    []byte
}

type SubscribeCallbackFun func(message *Message)

type Subscriber interface {
	Subscribe(subject string, callback SubscribeCallbackFun) (subscription Subscription, err error)
}

type Publisher interface {
	Publish(subject string, message *Message) (err error)
}

type Subscription interface {
	Unsubscribe() error
}
