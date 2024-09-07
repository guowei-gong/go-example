# 泛型

为什么 Go 需要泛型？因为泛型可以**约束入参类型**，any 或 interface 类型不可以。

为什么需要约束入参类型？因为约束类型，可以**避免许多类型转换的问题**。

## 语法
泛型接口，T 称作参数名称，any 称作类型约束。下面的例子中，参数 T 可以是任意类型。
```go
type ICache[T any] interface {
    // Set 保存数据
    Set(id string, val T)
    // Del 删除数据
    Del(id string)
    // Get 获取数据
    Get(id string) (T, error)
}
```

结构体，也可以使用泛型。
```go
type RedisCache[T any] struct {
    store map[string]T
}
```

方法，也可以使用泛型。


## 练习
用泛型实现以下函数。

- 求和；
- 过滤和查找；
- 求最大值、最小值；
- 在切片特定索引位置插入元素。

