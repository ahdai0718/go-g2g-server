package logger

import "ohdada/g2gserver/internal/pkg/pb"

var (
	defaultSimpleFactory = &SimpleFactory{}
)

// DefaultSimpleFactory .
func DefaultSimpleFactory() *SimpleFactory {
	return defaultSimpleFactory
}

// Logger .
type Logger interface {
	SetLogServerInfo(serverInfo *pb.ServerInfo)
	Log(path string, data interface{}) error
}
