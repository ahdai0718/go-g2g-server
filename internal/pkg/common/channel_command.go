package common

// NewChannelCommand .
func NewChannelCommand(command int, data ...interface{}) *ChannelCommand {
	return &ChannelCommand{
		Command: command,
		Data:    data,
	}
}

// NewSyncChannelCommand .
func NewSyncChannelCommand(command int, data ...interface{}) *ChannelCommand {
	return &ChannelCommand{
		Command: command,
		Data:    data,
		Channel: make(chan interface{}),
	}
}

// ChannelCommand .
type ChannelCommand struct {
	Command int
	Data    []interface{}
	Channel chan interface{}
}

// ChannelCommandHandler .
type ChannelCommandHandler interface {
	Handle(*ChannelCommand)
}

// ChannelCommandHandlerFactory .
type ChannelCommandHandlerFactory interface {
	Create(*ChannelCommand, ...interface{}) ChannelCommandHandler
}
