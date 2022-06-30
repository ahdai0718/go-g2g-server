package cchandler

import (
	"ohdada/g2gserver/internal/pkg/common"
)

// PlayerChannelCommand
const (
	_ = iota
)

var (
	logInfoMap = map[int]string{}
)

// SimpleFactory .
type SimpleFactory struct{}

// Create .
func (factory *SimpleFactory) Create(command *common.ChannelCommand, subjectList ...interface{}) (handler common.ChannelCommandHandler) {

	return
}
