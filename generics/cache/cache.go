package cache

import "errors"

type ICache[T any] interface {
	// Set 保存数据
	Set(id string, val T)
	// Del 删除数据
	Del(id string)
	// Get 获取数据
	Get(id string) (T, error)
}

type RedisCache[T any] struct {
	store map[string]T
}

// NewRedisCache 创建一个新的 RedisCache 实例
func NewRedisCache[T any]() *RedisCache[T] {
	return &RedisCache[T]{store: make(map[string]T)}
}

func (r *RedisCache[T]) Set(id string, val T) {
	r.store[id] = val
}

func (r *RedisCache[T]) Del(id string) {
	delete(r.store, id)
}

func (r *RedisCache[T]) Get(id string) (T, error) {
	v, ok := r.store[id]
	if !ok {
		var notFound T
		return notFound, errors.New("not found")
	}
	return v, nil
}
