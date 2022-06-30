package gateway

import (
	"ohdada/g2gserver/internal/pkg/config"
	"ohdada/g2gserver/internal/pkg/constant/envname"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/mq"
)

var (
	messager mq.Messager
)

func initMQ() {

	mqType := config.GetString(envname.MQType)
	mqUris := config.GetString(envname.MQUris)

	messager = mq.DefaultSimpleFactory().Create(mqType)

	if err := messager.Init(mqUris); err != nil {
		glog.Error(err)
	}

	if err := messager.Run(); err != nil {
		glog.Error(err)
	}

	// messager.AddChannel("Common", "Common.*")

	_, err := messager.Subscribe("Test.*", func(message *mq.Message) {
		glog.Infoln(string(message.Data))
	})

	if err != nil {
		glog.Error(err)
	}

	messager.Publish("Test.test", &mq.Message{
		Data: []byte("Hello, i am gateway server"),
	})
}

func closeMQ() {
	if err := messager.Close(); err != nil {
		glog.Error(err)
	}
}
