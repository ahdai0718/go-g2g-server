package logger

import "github.com/go-resty/resty/v2"

// logger log type
const (
	_ = iota
	TypeBigQuery
)

// SimpleFactory .
type SimpleFactory struct{}

// Create .
func (factory *SimpleFactory) Create(logType int64) (logger Logger) {

	switch logType {
	case TypeBigQuery:
		logger = &BigQuery{
			client: resty.New(),
		}
	default:
		logger = &Default{}
	}

	return
}
