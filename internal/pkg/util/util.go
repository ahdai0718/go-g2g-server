package util

import (
	"math/rand"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/runtime/protoiface"
)

// Rand .
func Rand(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	if max-min > 0 {
		return rand.Intn(max-min) + min
	}

	return min
}

// Rand64 .
func Rand64(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())

	if max-min > 0 {
		return rand.Int63n(max-min) + min
	}

	return min
}

// Timestamp .
func Timestamp() int64 {
	return time.Now().Unix()
}

// UUID .
func UUID() string {
	return uuid.New().String()
}

// AbsInt32 .
func AbsInt32(x int32) int32 {
	if x < 0 {
		return -x
	}
	return x
}

// Int32ToPercentage ...
func Int32ToPercentage(x int32) float64 {
	return float64(x) / float64(100)
}

// IsSameDay .
func IsSameDay(time1 time.Time, time2 time.Time) (isSameDay bool) {

	isSameDay = false

	if time1.Day() == time2.Day() &&
		time1.Month() == time2.Month() &&
		time1.Year() == time2.Year() {
		isSameDay = true
	}

	return
}

// ConvertProtoMessageToJSONString .
func ConvertProtoMessageToJSONString(message protoiface.MessageV1) (jsonString string, err error) {
	marshaler := jsonpb.Marshaler{}
	marshaler.EmitDefaults = true
	marshaler.OrigName = true
	marshaler.EnumsAsInts = true
	jsonString, err = marshaler.MarshalToString(message)
	return
}
