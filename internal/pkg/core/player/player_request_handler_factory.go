package player

import (
	"ohdada/g2gserver/internal/pkg/common"
	"ohdada/g2gserver/internal/pkg/pb"
)

// RequestHandlerSimpleFactory .
type RequestHandlerSimpleFactory interface {
	CreateRequestHandler(*pb.Request) common.RequestHandler
}
