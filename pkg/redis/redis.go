package redis

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/literaen/simple_project/pkg/config"

	"github.com/redis/go-redis/v9"
)

type RDB struct {
	client *redis.Client
}

func NewRDB(creds *config.REDIS_CREDS) *RDB {
	dsn := fmt.Sprintf("%s:%s", creds.Host, creds.Port)
	client := redis.NewClient(&redis.Options{
		Addr: dsn,
		DB:   0,
	})

	res, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("failed to ping redis %v", err)
	} else {
		log.Printf("redis ping: %v", res)
	}

	return &RDB{client: client}
}

func (r *RDB) Set(ctx context.Context, key string, value string, ttl int) error {
	return r.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

func (r *RDB) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}

func (r *RDB) Del(ctx context.Context, key string) error {
	err := r.client.Del(ctx, key).Err()

	return err
}

func (r *RDB) Close() error {
	return r.client.Close()
}

func (r *RDB) HSet(ctx context.Context, key string, value interface{}) error {
	err := r.client.HSet(ctx, key, value).Err()

	return err
}

func (r *RDB) HGet(ctx context.Context, key, field string) (string, error) {
	res, err := r.client.HGet(ctx, key, field).Result()
	if errors.Is(err, redis.Nil) {
		return "", nil
	}

	return res, err
}

func (r *RDB) HDel(ctx context.Context, key string) error {
	return r.client.HDel(ctx, key).Err()
}

func (r *RDB) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	return r.client.HGetAll(ctx, key).Result()
}

func (r *RDB) RPush(ctx context.Context, key string, value interface{}) (err error) {
	err = r.client.RPush(ctx, key, value).Err()

	return err
}

func (r *RDB) LRange(ctx context.Context, key string, start, stop int64) (err error) {
	err = r.client.LRange(ctx, key, start, stop).Err()

	return err
}

func (r *RDB) Ping(ctx context.Context) error {
	err := r.client.Ping(ctx).Err()

	return err
}
