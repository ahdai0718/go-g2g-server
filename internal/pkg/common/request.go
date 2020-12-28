package common

import "ohdada/g2gserver/internal/pkg/pb"

// RequestHandler .
type RequestHandler interface {
	Handle(*pb.Request)
}
