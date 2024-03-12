package redis_pkg

import (
	"context"

	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"

	redis_store "github.com/eko/gocache/store/redis/v4"
)



func RedisClient() (*redis_store.RedisStore, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	// Add tracing to all commands
	err = redisotel.InstrumentTracing(rdb)
	if err != nil {
		return nil, err
	}

	redisStore := redis_store.NewRedis(rdb)
	return redisStore, nil
}
