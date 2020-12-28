package room

import (
	"github.com/jinzhu/copier"
)

// NewBuilder .
func NewBuilder() *Builder {
	return &Builder{}
}

// Builder .
type Builder struct {
	ccHandlerBufferSize    int
	ccHandlerSimpleFactory CCHandlerSimpleFactory
}

// Build .
func (builder *Builder) Build() *Room {

	room := &Room{}

	copier.Copy(room, builder)

	return room
}
