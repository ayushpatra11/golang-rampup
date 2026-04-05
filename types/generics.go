package types

import "fmt"

// Generics added in Go 1.18 (type parameters)
// Syntax: func Name[T constraint](args) returnType

// any = interface{} -> no constraint (any type)
func Map[T, U any](slice []T, f func(T) U) []U {
	result := make([]U, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func Filter[T any](slice []T, pred func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U {
	acc := initial
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}

// comparable: built-in constraint for types supporting == and !=
func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// custom constraint: union of types using |
// ~ means "underlying type" - allows type aliases
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// generic data structures
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int { return len(s.items) }

func genericsDemo() {
	nums := []int{1, 2, 3, 4, 5}

	doubled := Map(nums, func(n int) int { return n * 2 })
	fmt.Println(doubled) // [2 4 6 8 10]

	evens := Filter(nums, func(n int) bool { return n%2 == 0 })
	fmt.Println(evens) // [2 4]

	total := Reduce(nums, 0, func(acc, n int) int { return acc + n })
	fmt.Println(total) // 15

	fmt.Println(Sum(nums))                  // 15
	fmt.Println(Contains(nums, 3))          // true
	fmt.Println(Contains([]string{"a", "b"}, "c")) // false

	var s Stack[string]
	s.Push("hello")
	s.Push("world")
	v, _ := s.Pop()
	fmt.Println(v) // world
}

/*
Type inference: Go can often infer type parameters from the arguments.
  Sum[int](nums) -> can be written as Sum(nums)
  Map[int, string](nums, f) -> can be written as Map(nums, f)

When to use generics:
- Operating on containers/collections uniformly across types
- Avoiding code duplication that type aliases can't fix
Don't use for behavior that varies per type -> use interfaces instead.
*/
