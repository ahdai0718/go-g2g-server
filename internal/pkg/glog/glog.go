package glog

import (
	"encoding/json"
	"fmt"
	"net/url"
	"ohdada/g2gserver/internal/pkg/pb"
	"os"
	"runtime"
	"time"

	"github.com/go-resty/resty/v2"
	googlelog "github.com/golang/glog"
	"github.com/sirupsen/logrus"
)

var (
	logServerInfo     *pb.ServerInfo
	logRoutePath      = "log.runtime_log"
	currentServerInfo *pb.ServerInfo
	restyClient       = resty.New()
	isSendLogToServer = true
	log               = logrus.New()
)

func Init() {
	log.SetOutput(os.Stdout)
	// log.SetFormatter(&logrus.JSONFormatter{})
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors: true,
	})
}

func caller() (file string, line int) {
	_, file, line, ok := runtime.Caller(2)

	if !ok {
		file = "???"
		line = -1
	} else {

		// slash := strings.LastIndex(file, "/")
		// if slash >= 0 {
		// 	file = file[slash+1:]
		// }

	}

	return file, line
}

func sendRuntimeLogToLogServer(runtimeLog *pb.RuntimeLog) (err error) {

	if logServerInfo == nil {
		return
	}

	data, err := json.Marshal(runtimeLog)
	if err != nil {
		googlelog.Error(err)
		return
	}

	body := make(map[string]string)
	body["json"] = string(data)

	url := &url.URL{
		Scheme: logServerInfo.Protocol,
		Host:   fmt.Sprintf("%s:%d", logServerInfo.Host, logServerInfo.Port),
		Path:   logRoutePath,
	}

	_, err = restyClient.R().SetFormData(body).Post(url.String())
	if err != nil {
		googlelog.Error(err)
		return
	}

	return
}

func sendLogToServer(logType pb.RuntimeLogType, file string, line int, args interface{}) {
	if isSendLogToServer {
		sendRuntimeLogToLogServer(&pb.RuntimeLog{
			Type:       int64(logType),
			File:       file,
			Line:       int64(line),
			Message:    fmt.Sprintf("%s", args),
			ServerInfo: currentServerInfo,
			Time:       getTimeNowUnixMicro(),
		})
	}
}

func getTimeNowUnixMicro() int64 {
	return time.Now().UnixNano() / 1e3
}

// TurnLogToServerOn .
func TurnLogToServerOn() {
	isSendLogToServer = true
}

// TurnLogToServerOff .
func TurnLogToServerOff() {
	isSendLogToServer = false
}

// SetCurrentServerInfo .
func SetCurrentServerInfo(serverInfo *pb.ServerInfo) {
	currentServerInfo = serverInfo
}

// SetLogServerInfo .
func SetLogServerInfo(serverInfo *pb.ServerInfo) {
	logServerInfo = serverInfo
}

// SetLogRoutePath .
func SetLogRoutePath(routePath string) {
	logRoutePath = routePath
}

// V .
func V(level googlelog.Level) googlelog.Verbose {
	return googlelog.V(level)
}

// Info .
func Info(args ...interface{}) {

	file, line := caller()

	// googlelog.Info(append([]interface{}{file, line}, args...))

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Info(args...)
}

// Infof .
func Infof(format string, args ...interface{}) {

	file, line := caller()

	// googlelog.Info(file, ":", line)

	// googlelog.Infof(format, args...)

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Infof(format, args...)
}

// Infoln .
func Infoln(args ...interface{}) {

	file, line := caller()

	// googlelog.Infoln(append([]interface{}{file, line}, args...))

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Infoln(args...)

}

// Warning .
func Warning(args ...interface{}) {

	file, line := caller()

	// googlelog.Warning(append([]interface{}{file, line}, args...))

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Warning(args...)
}

// Warningf .
func Warningf(format string, args ...interface{}) {

	file, line := caller()

	// googlelog.Warning(file, ":", line)

	// googlelog.Warningf(format, args...)

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Warningf(format, args...)

}

// Warningln .
func Warningln(args ...interface{}) {

	file, line := caller()

	// googlelog.Warningln(append([]interface{}{file, line}, args...))

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Warningln(args...)
}

// Error .
func Error(args ...interface{}) {

	file, line := caller()

	// googlelog.Error(append([]interface{}{file, line}, args...))

	sendLogToServer(pb.RuntimeLogType_RLT_ERROR, file, line, args)

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Error(args...)
}

// Errorf .
func Errorf(format string, args ...interface{}) {

	file, line := caller()

	// googlelog.Error(file, ":", line)

	// googlelog.Errorf(format, args...)

	sendLogToServer(pb.RuntimeLogType_RLT_ERROR, file, line, fmt.Sprintf(format, args...))

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Errorf(format, args...)
}

// Errorln .
func Errorln(args ...interface{}) {

	file, line := caller()

	// googlelog.Errorln(append([]interface{}{file, line}, args...))

	sendLogToServer(pb.RuntimeLogType_RLT_ERROR, file, line, args)

	log.WithFields(logrus.Fields{
		"file": file,
		"line": line,
	}).Errorln(args...)
}
