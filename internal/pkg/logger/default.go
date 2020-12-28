package logger

import "ohdada/g2gserver/internal/pkg/pb"

// Default .
type Default struct{}

// SetLogServerInfo .
func (logger *Default) SetLogServerInfo(serverInfo *pb.ServerInfo) {}

// Log .
func (logger *Default) Log(path string, data interface{}) (err error) { return }

// LogPlatform .
func (logger *Default) LogPlatform(platformLog *pb.PlatformLog) (err error) { return }
