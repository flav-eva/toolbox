package redis

import (
	"crypto/tls"
	"sync"

	"github.com/go-redis/redis/v8"
)

var db *redis.Client
var once sync.Once

func Init(addr, password string, database int, enableTls bool) {
	once.Do(func() {
		opts := &redis.Options{
			Addr:     addr,
			Password: password,
			DB:       database,
		}
		if enableTls {
			opts.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		}

		db = redis.NewClient(opts)
	})
}

func Client() *redis.Client {
	if db == nil {
		panic("redis client not init")
	}
	return db
}
