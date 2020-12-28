package logger

import (
	"ohdada/g2gserver/internal/pkg/pb"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

// BigQuery .
type BigQuery struct {
	serverInfo *pb.ServerInfo
	client     *resty.Client
}

// SetLogServerInfo .
func (logger *BigQuery) SetLogServerInfo(serverInfo *pb.ServerInfo) {
	logger.serverInfo = serverInfo
}

// Log .
func (logger *BigQuery) Log(path string, data interface{}) (err error) {

	url := &url.URL{
		Scheme: logger.serverInfo.Protocol,
		Host:   fmt.Sprintf("%s:%d", logger.serverInfo.Host, logger.serverInfo.Port),
		Path:   path,
	}

	return logger.postData(url.String(), data)
}

func (logger *BigQuery) postData(url string, data interface{}) error {

	dataByteArray, err := json.Marshal(data)

	if err != nil {
		return err
	}

	dataMap := make(map[string]string)
	dataMap["json"] = string(dataByteArray)

	_, err = logger.client.R().SetFormData(dataMap).Post(url)

	return err
}
