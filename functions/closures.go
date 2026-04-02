package functions

import "fmt"

// functions are first-class values in Go - can be assigned, passed, returned

// function type
type transformer func(int) int

func applyTwice(f transformer, x int) int {
	return f(f(x))
}

func applyTwiceDemo() {
	double := func(x int) int { return x * 2 }
	fmt.Println(applyTwice(double, 3)) // 12
}

// closure: a function that captures variables from its surrounding scope
// the captured variable is shared (not copied) - mutations are visible
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func closureDemo() {
	c1 := makeCounter()
	c2 := makeCounter() // independent counter; own captured `count`

	fmt.Println(c1(), c1(), c1()) // 1 2 3
	fmt.Println(c2())             // 1
}

// factory pattern with closures
func makeMultiplier(factor int) func(int) int {
	return func(x int) int { return x * factor }
}

// memoization using closure over a map
func memoize(f func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		result := f(n)
		cache[n] = result
		return result
	}
}

// classic gotcha: closure over loop variable
func closureLoopBug() []func() int {
	funcs := make([]func() int, 3)
	for i := 0; i < 3; i++ {
		i := i // shadow with a new variable each iteration - fixes the bug
		funcs[i] = func() int { return i }
	}
	return funcs
}
// Without the `i := i` line all closures would return 3 (loop final value)

// immediately invoked function expression (IIFE)
func iifeDemo() {
	result := func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println(result) // 7
}

// variadic function: accepts zero or more arguments of a type
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func variadicDemo() {
	fmt.Println(sum(1, 2, 3))    // 6
	fmt.Println(sum())            // 0

	s := []int{4, 5, 6}
	fmt.Println(sum(s...))        // 15 - unpack slice with ...
}

/*
Closures capture variables by reference.
If you need to capture a value (snapshot), either:
  1. Shadow the variable: x := x
  2. Pass it as a function argument to the closure
*/
