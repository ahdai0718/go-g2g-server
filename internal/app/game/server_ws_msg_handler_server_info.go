package game

import (
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/pb"

	"google.golang.org/protobuf/proto"
)

// ServerWSMessageHandlerServerInfo .
type ServerWSMessageHandlerServerInfo struct {
}

// Handle .
func (handler *ServerWSMessageHandlerServerInfo) Handle(message []byte) {

	serverInfo := &pb.ServerInfo{}

	err := proto.Unmarshal(message, serverInfo)

	if err != nil {
		glog.Error(err)
		return
	}

	glog.Infoln("ServerWSMessageHandlerServerInfo:", serverInfo)

}
