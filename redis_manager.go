package bzredis

import (
	"context"
	"errors"
)

type MapStringOptions map[string]*Options

var mgr = map[string]*Redis{}

func Init(mo MapStringOptions) {
	for key, opt := range mo {
		mgr[key] = NewRedis(opt)
	}
}

func InitAndPing(ctx context.Context, mo MapStringOptions) {
	for key, opt := range mo {
		m, err := NewRedisPing(ctx, opt)
		if err != nil {
			continue
		}

		mgr[key] = m
	}
}

func GetRedis(name string) (*Redis, error) {
	m, ok := mgr[name]
	if !ok {
		return nil, errors.New("No redis named:" + name)
	}

	return m, nil
}
