package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang/glog"
)

type clientRedisCluster struct {
	clientBase
	client *redis.ClusterClient
}

// Init .
func (client *clientRedisCluster) Init(serverConfigList []*ServerConfig) error {

	if len(serverConfigList) == 0 {
		return fmt.Errorf("sever config is empty")
	}

	clusterOptions := &redis.ClusterOptions{}

	for _, serverConfig := range serverConfigList {
		clusterOptions.Addrs = append(clusterOptions.Addrs, fmt.Sprintf("%s:%d", serverConfig.Host, serverConfig.Port))
	}

	clusterOptions.Password = serverConfigList[0].Password

	client.client = redis.NewClusterClient(clusterOptions)

	statusCmd := client.client.Ping(context.Background())

	return statusCmd.Err()
}

// Set .
func (client *clientRedisCluster) Set(key string, value interface{}) error {
	statusCmd := client.client.Set(context.Background(), key, value, 0)
	return statusCmd.Err()
}

// SetWithExpiration .
func (client *clientRedisCluster) SetWithExpiration(expiration time.Duration, key string, value interface{}) error {
	statusCmd := client.client.Set(context.Background(), key, value, expiration)
	return statusCmd.Err()
}

// Get .
func (client *clientRedisCluster) Get(key string, value interface{}) error {
	statusCmd := client.client.Get(context.Background(), key)
	err := statusCmd.Scan(value)
	return err
}

// Delete .
func (client *clientRedisCluster) Delete(key string) error {
	statusCmd := client.client.Del(context.Background(), key)
	return statusCmd.Err()
}

// Push .
func (client *clientRedisCluster) Push(key string, values ...interface{}) error {
	statusCmd := client.client.RPush(context.Background(), key, values)
	return statusCmd.Err()
}

// PushWithExpiration .
func (client *clientRedisCluster) PushWithExpiration(expiration time.Duration, key string, values ...interface{}) error {
	statusCmd := client.client.RPush(context.Background(), key, values)
	client.client.Expire(context.Background(), key, expiration)
	return statusCmd.Err()
}

// Pop .
func (client *clientRedisCluster) Pop(key string, value interface{}) (err error) {
	statusCmd := client.client.LPop(context.Background(), key)
	err = statusCmd.Scan(value)

	if err != nil {
		glog.Error(err)
		return
	}

	err = statusCmd.Err()

	return
}

// PopAll .
func (client *clientRedisCluster) PopAll(key string, value interface{}) (err error) {

	len, err := client.client.LLen(context.Background(), key).Result()

	if err != nil {
		glog.Error(err)
		return
	}

	if len == 0 {
		return
	}

	statusCmd := client.client.LRange(context.Background(), key, 0, len-1)

	err = statusCmd.ScanSlice(value)

	if err != nil {
		glog.Error(err)
		return
	}

	err = statusCmd.Err()

	if err == nil {
		client.client.LTrim(context.Background(), key, len, -1)
	}

	return
}

// HashGet .
func (client *clientRedisCluster) HashGet(key string, field string, value interface{}) error {
	statusCmd := client.client.HGet(context.Background(), key, field)
	err := statusCmd.Scan(value)
	return err
}

// HashGetAll .
func (client *clientRedisCluster) HashGetAll(key string, value interface{}) (err error) {
	if _, isOk := value.(map[string]string); !isOk {
		err = fmt.Errorf("value should be map[string]string")
		return
	}
	stringMapCmd := client.client.HGetAll(context.Background(), key)
	result, err := stringMapCmd.Result()
	for k, v := range result {
		value.(map[string]string)[k] = v
	}
	return
}

// HashSet .
func (client *clientRedisCluster) HashSet(key string, field string, value interface{}) error {
	statusCmd := client.client.HSet(context.Background(), key, field, value)
	return statusCmd.Err()
}

// HashDelete .
func (client *clientRedisCluster) HashDelete(key string, field string) error {
	intCmd := client.client.HDel(context.Background(), key, field)
	return intCmd.Err()
}
