package bzredis

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisOptions interface {
	RedisOptions() *redis.Options
}

type Options struct {
	Addr     string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB       int    `json:"db" yaml:"db"`

	DialTimeout  time.Duration `json:"dial_timeout" yaml:"dial_timeout"`
	ReadTimeout  time.Duration `json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout time.Duration `json:"write_timeout" yaml:"write_timeout"`
	MaxRetries   int           `json:"max_retries" yaml:"max_retries"`
	PoolSize     int           `json:"pool_size" yaml:"pool_size"`
	MinIdleConns int           `json:"min_idle_conns" yaml:"min_idle_conns"`
}

// RedisOptions 转换为 redis.Options
func (c *Options) RedisOptions() *redis.Options {
	return &redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,

		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		MaxRetries:   c.MaxRetries,
		PoolSize:     c.PoolSize,
		MinIdleConns: c.MinIdleConns,
	}
}

func DefaultOptions() *Options {
	return &Options{
		Addr: "127.0.0.1:6379",
	}
}

func (o *Options) WithAddr(addr string) *Options {
	o.Addr = addr
	return o
}

func (o *Options) WithPassword(password string) *Options {
	o.Password = password
	return o
}
func (o *Options) WithDB(db int) *Options {
	o.DB = db
	return o
}
func (o *Options) WithDialTimeout(dialTimeout time.Duration) *Options {
	o.DialTimeout = dialTimeout
	return o
}
func (o *Options) WithReadTimeout(readTimeout time.Duration) *Options {
	o.ReadTimeout = readTimeout
	return o
}
func (o *Options) WithWriteTimeout(writeTimeout time.Duration) *Options {
	o.WriteTimeout = writeTimeout
	return o
}
func (o *Options) WithMaxRetries(maxRetries int) *Options {
	o.MaxRetries = maxRetries
	return o
}
func (o *Options) WithPoolSize(poolSize int) *Options {
	o.PoolSize = poolSize
	return o
}
func (o *Options) WithMinIdleConns(minIdleConns int) *Options {
	o.MinIdleConns = minIdleConns
	return o
}
