package gateway

import (
	"ohdada/g2gserver/internal/pkg/core/serverinfo"
	"ohdada/g2gserver/internal/pkg/glog"
	"ohdada/g2gserver/internal/pkg/network"
	"ohdada/g2gserver/internal/pkg/pb"

	"google.golang.org/protobuf/proto"
)

var (
	serverWebsocketConnectionManager *ServerWebsocketConnectionManager
)

func defaultServerWebsocketConnectionManager() *ServerWebsocketConnectionManager {
	if serverWebsocketConnectionManager == nil {
		serverWebsocketConnectionManager = &ServerWebsocketConnectionManager{
			addWebsocketConnectionChannel: make(chan *network.WebsocketConnection, 32),
		}
		serverWebsocketConnectionManager.Init()
	}
	return serverWebsocketConnectionManager
}

// ServerWebsocketConnectionManager .
type ServerWebsocketConnectionManager struct {
	network.WebsocketConnectionManager
	addWebsocketConnectionChannel chan *network.WebsocketConnection
}

// OnConnect .
func (manager *ServerWebsocketConnectionManager) OnConnect(wc *network.WebsocketConnection) {
	glog.Infoln("ServerWebsocketConnectionManager OnConnect", wc)
	manager.addWebsocketConnectionChannel <- wc
}

func (manager *ServerWebsocketConnectionManager) observeWC(wc *network.WebsocketConnection) {
	for {
		select {
		case request := <-wc.ReceiveRequestChannel():

			command := pb.RequestCommand(request.Command)

			handler := defaultServerWSMessaageHandlerSimpleFactory().create(command)
			handler.Handle(request.Data)

			if command == pb.RequestCommand_RC_SERVER_INFO {
				manager.sendServerInfo(wc)
			}

		case <-wc.CloseChannel():
			manager.RemoveWebsocketConnection(wc)
			serverinfo.DefaultManager().RemoveRunningServerByName(wc.ID())
			return
		}
	}
}

func (manager *ServerWebsocketConnectionManager) observe() {
	for {
		select {
		case wc := <-manager.addWebsocketConnectionChannel:
			if !manager.IsWebsocketConnectionExists(wc) {
				manager.AddWebsocketConnection(wc)
				go manager.observeWC(wc)
			}
		}
	}
}

func (manager *ServerWebsocketConnectionManager) sendServerInfo(wc *network.WebsocketConnection) {
	serverInfo := ServerInfo()

	dataServerInfo, err := proto.Marshal(serverInfo)

	if err != nil {
		glog.Error(err)
		return
	}

	wc.SendRequest(int32(pb.RequestCommand_RC_SERVER_INFO), dataServerInfo)
}
