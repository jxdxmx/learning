package backup

import (
	"fmt"
	"sort"
)

// Ordered 类型限制
type Ordered interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}

type orderedSlice[T Ordered] []T

func (s orderedSlice[T]) Len() int           { return len(s) }
func (s orderedSlice[T]) Less(i, j int) bool { return s[i] < s[j] }
func (s orderedSlice[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type Node[T Ordered] struct {
	a, b T
}

func F20220706() {
	var s orderedSlice[byte]
	s = append(s, 5, 1, 2, 7)
	sort.Sort(s)
	fmt.Println(s)

	var node Node[int]
	node.a = 60
	fmt.Println(node.b + node.b)
}
