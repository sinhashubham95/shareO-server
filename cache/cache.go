package cache

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	"log"
	"time"
)

var (
	client      *redis.Client
	errNoClient = errors.New("redis not running")
)

func init() {
	url := "localhost:6379"
	client = redis.NewClient(&redis.Options{
		Addr:         url,
		PoolSize:     20,
		MinIdleConns: 10,
	})
	res, err := client.Ping().Result()
	if err != nil {
		Close()
		log.Fatalf("Test redis connection %+v.", res)
	}
}

// GetInterface returns the value in cache for the key passes as parameter
func GetInterface(key string, dto interface{}) error {
	if client != nil {
		res, err := client.Get(key).Bytes()
		if err != nil {
			return err
		}
		err = json.Unmarshal(res, dto)
		if err != nil {
			return err
		}
		return nil
	}
	return errNoClient
}

// SetInterface saves the value for key with a timeout of 10 days
func SetInterface(key string, val interface{}) error {
	return SetWithExpiration(key, val, 240*time.Hour, true)
}

// SetWithExpiration saves the value for key and val in cache
func SetWithExpiration(key string, val interface{}, expiration time.Duration, parse bool) error {
	var res interface{}
	var err error
	if parse {
		res, err = json.Marshal(val)
		if err != nil {
			return err
		}
	} else {
		res = val
	}
	if client != nil {
		err = client.Set(key, res, expiration).Err()
		if err != nil {
			return err
		}
		return nil
	}
	return errNoClient
}

// Close closes the redis connection
func Close() {
	if client == nil {
		return
	}
	client.Close()
	client = nil
}
