package bzredis

import (
	"context"
	"testing"
)

func TestRedis(t *testing.T) {
	ctx := context.Background()
	r, err := NewRedisPing(ctx, DefaultOptions())
	if err != nil {
		t.Fatal(err)
	}

	r.Set(ctx, "test_name", 100, 0)
	v, _ := r.Get(ctx, "test_name").Result()

	t.Log(v)
}
