package main

import (
	"fmt"
	"github/go-example/generics/cache"
)

const (
	key = "1"
)

// 假设使用 RedisCache
func main() {
	c := cache.NewRedisCache[int]()

	c.Set(key, 1024)
	// 保存字符串类型，编译器会报错，因为变量 c 约束保存的值为 int 类型
	// 假如 ICache 是 Set(id string, val any)，编译器不会报错，也就没有约束类型
	// c.Set(key, "1024")

	val, err := c.Get(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(val)

	c.Del(key)
}
