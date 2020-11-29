package bzredis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	*redis.Client
}

func NewRedis(opt RedisOptions) *Redis {
	cli := redis.NewClient(opt.RedisOptions())

	return &Redis{Client: cli}
}

func NewRedisPing(ctx context.Context, opt RedisOptions) (*Redis, error) {
	cli := redis.NewClient(opt.RedisOptions())

	err := cli.Ping(context.Background()).Err()
	if err != nil {
		return nil, err
	}

	return &Redis{
		Client: cli,
	}, nil
}
