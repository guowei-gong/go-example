package practise

import (
	"fmt"
	"github.com/ecodeclub/ekit"
)

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Integer int

func main() {
	// 求和
	fmt.Println("Sum:", Sum[int](1, 1))
	// Integer 是 int 的类型别名，也可作为泛型的衍生类型
	fmt.Println("Sum:", Sum[Integer](1, 1))

	// 过滤和查找

	// 求最大值、最小值
	fmt.Println("Max:", Max[float32](1.1, 1.2))
	fmt.Println("Min:", Min[float64](1.1, 1.2))

	// 在切片特定索引位置插入元素。
}

func Sum[T Number](data ...T) (res T) {
	for _, v := range data {
		res += v
	}
	return
}

// Contains 判断 src 里面是否存在 dst
func Contains[T comparable](src []T, dst T) bool {
	return ContainsFunc[T](src, func(src T) bool {
		return src == dst
	})
}

// ContainsFunc 优先使用 Contains
func ContainsFunc[T any](src []T, equal func(src T) bool) bool {
	for _, v := range src {
		if equal(v) {
			return true
		}
	}
	return false
}

func Filter[T comparable](src []T, dst T) (res []T) {
	for _, v := range src {
		if v == dst {
			continue
		}
		res = append(res, v)
	}
	return
}

func FilterMap[Src any, Dst any](src []Src, filter func(idx int, src Src) (Dst, bool)) []Dst {
	res := make([]Dst, 0, len(src))
	for i, s := range src {
		dst, ok := filter(i, s)
		if ok {
			res = append(res, dst)
		}
	}
	return res
}

func Max[T ekit.RealNumber](data ...T) (res T) {
	res = data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > res {
			res = data[i]
		}
	}
	return res
}

func Min[T ekit.RealNumber](data ...T) (res T) {
	res = data[0]
	for i := 1; i < len(data); i++ {
		if data[i] < res {
			res = data[i]
		}
	}
	return
}
