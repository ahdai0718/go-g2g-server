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
	ccHandlerSimpleFactory CCHandlerSimpleFactory
	ccHandlerBufferSize    int
}

// Build .
func (builder *Builder) Build() *Room {

	room := &Room{}

	copier.Copy(room, builder)

	return room
}
