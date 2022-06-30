package gateway

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"

	"google.golang.org/protobuf/proto"
)

// ServerWSMessageHandlerBroadcast .
type ServerWSMessageHandlerBroadcast struct {
}

// Handle .
func (handler *ServerWSMessageHandlerBroadcast) Handle(message []byte) {

	request := &pb.Request{}

	err := proto.Unmarshal(message, request)

	if err != nil {
		glog.Error(err)
		return
	}

	DefaultPlayerManager().broadcastInAsync(message)
}
