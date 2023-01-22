package test

import (
	"context"
	"github.com/go-redis/redis/v9"
	"light-cloud/src/core/define"
	"testing"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     define.RedisAddress,
	Password: define.RedisPassword,
	DB:       0, // use default DB
})

func TestSetValue(t *testing.T) {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Error(err)
	}
}

func TestGetValue(t *testing.T) {
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Error(err)
	}
	t.Log(val)
}
